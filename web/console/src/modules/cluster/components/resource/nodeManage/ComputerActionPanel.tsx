import { connect } from 'react-redux';
import { allActions } from '../../../actions';
import * as React from 'react';
import { bindActionCreators } from '@tencent/qcloud-lib';
import { Button, SearchBox, Bubble, Modal, TabPanel, Table, Text } from '@tea/component';
import { Computer } from '../../../models';
import { RootProps } from '../../ClusterApp';
import { downloadCsv, dateFormatter } from '../../../../../../helpers';
import { FetchState } from '@tencent/qcloud-redux-fetcher';
import { MonitorPanelProps, nodeMonitorFields, podMonitorFields } from '../../../models/MonitorPanel';
import { t, Trans } from '@tencent/tea-app/lib/i18n';
import { Justify } from '@tea/component/justify';
import { ChartInstancesPanel } from '@tencent/tchart';
import { router } from '../../../../../modules/cluster/router';
import { ReduceRequest } from '../resourceDetail/ResourcePodPanel';
import { WorkflowDialog } from 'src/modules/common';

interface State {
  monitorPanelProps?: MonitorPanelProps;
  computerById?: any;
  showOsTips?: boolean;
  selectCluster?: any;
}
const mapDispatchToProps = dispatch =>
  Object.assign({}, bindActionCreators({ actions: allActions }, dispatch), { dispatch });

@connect(state => state, mapDispatchToProps)
export class ComputerActionPanel extends React.Component<RootProps, State> {
  state = {
    monitorPanelProps: null,
    showOsTips: false,
    selectCluster: null
  };
  componentDidMount() {
    let { actions, route } = this.props;
    let { rid, clusterId } = route.queries;
    actions.computer.poll({ clusterId, regionId: +rid });
  }
  componentWillUnmount() {
    let { actions } = this.props;
    actions.computer.performSearch('');
    actions.computer.clearPolling();
  }

  downloadHandle(computers: Computer[]) {
    let rows = [],
      head = ['ID', t('状态'), t('角色'), t('配置'), t('ip地址'), t('PodCIDR'), t('创建时间')];

    computers.forEach((item: Computer) => {
      let row = [
        item.metadata.name,
        item.status.phase,
        item.metadata.role,
        this._reduceCapacity(item),
        this._reduceIp(item),
        item.spec.podCIDR,
        dateFormatter(new Date(item.metadata.creationTimestamp), 'YYYY-MM-DD HH:mm:ss')
      ];
      rows.push(row);
    });

    downloadCsv(rows, head, `tke-computers${Date.now()}.csv`);
  }

  _reduceCapacity(node: Computer) {
    let capacity = node.status.capacity;
    let capacityInfo = {
      cpu: capacity.cpu,
      memory: capacity.memory
    };
    let finalCpu = ReduceRequest('cpu', capacityInfo),
      finalmem = (ReduceRequest('memory', capacity) / 1024).toFixed(2);
    return finalCpu + '核,' + finalmem + 'GB';
  }

  _reduceIp(node: Computer) {
    let finalIPInfo = node.status.addresses.filter(item => item.type !== 'Hostname');
    let ips = finalIPInfo.map((item, index) => item.address);
    return ips.join('、');
  }

  /** 处理 封锁 和 取消封锁的按钮的信息提示 */
  handleNodeErrTips(nodeArr) {
    if (nodeArr.length <= 3) {
      return nodeArr.map(item => item.instanceId).join('、');
    } else {
      return (
        '' +
        nodeArr[0].instanceId +
        '、' +
        nodeArr[1].instanceId +
        '、......、' +
        nodeArr[nodeArr.length - 1].instanceId
      );
    }
  }

  render() {
    const { actions, subRoot, route, cluster } = this.props,
      { computer } = subRoot.computerState;

    const disableAddNode = cluster.selection && cluster.selection.spec.type === 'Imported';

    let monitorButton = null;
    monitorButton = (
      <Button
        type="primary"
        disabled={!computer.list.data.records.length}
        onClick={() => {
          if (!computer.list.data.records.length) {
            return;
          }
          this._handleMonitor('nodeMonitor', '');
        }}
      >
        {t('监控')}
      </Button>
    );
    return (
      <Table.ActionPanel>
        <Justify
          left={
            <React.Fragment>
              {!disableAddNode && (
                <Button
                  type="primary"
                  onClick={() =>
                    router.navigate(
                      { sub: 'sub', mode: 'create', type: 'nodeManange', resourceName: 'node' },
                      route.queries
                    )
                  }
                >
                  {t('添加节点')}
                </Button>
              )}
              {monitorButton}
            </React.Fragment>
          }
          right={
            <React.Fragment>
              <SearchBox
                value={computer.query.keyword || ''}
                onChange={actions.computer.changeKeyword}
                onSearch={actions.computer.performSearch}
                placeholder={t('请输入节点名')}
                onClear={() => {
                  actions.computer.performSearch('');
                }}
              />
              <Button
                icon="download"
                onClick={() => {
                  this.downloadHandle(computer.list.data.records);
                }}
                disabled={computer.list.loading || computer.list.fetchState === FetchState.Fetching}
              >
                {t('导出全部')}
              </Button>
            </React.Fragment>
          }
        />
        {this.state && this.state.monitorPanelProps && (
          <Modal
            visible={true}
            caption={this.state.monitorPanelProps.title}
            onClose={() => this.setState({ monitorPanelProps: undefined })}
            size={1050}
          >
            <Modal.Body>
              <ChartInstancesPanel
                tables={this.state.monitorPanelProps.tables}
                groupBy={this.state.monitorPanelProps.groupBy}
                instance={this.state.monitorPanelProps.instance}
              >
                {this.state.monitorPanelProps.headerExtraDOM}
              </ChartInstancesPanel>
            </Modal.Body>
          </Modal>
        )}
      </Table.ActionPanel>
    );
  }
  /** 处理监控的相关操作 */
  private _handleMonitor(monitorType?: string, nodeName?: string) {
    let { computerState } = this.props.subRoot;

    let monitorPanelProps =
      monitorType === 'nodeMonitor'
        ? {
            tables: [
              {
                table: 'k8s_node',
                conditions: [
                  ['tke_cluster_instance_id', '=', this.props.route.queries.clusterId],
                  ['node_role', '=', 'Node']
                ],
                fields: nodeMonitorFields
              }
            ],
            groupBy: [{ value: 'node' }],
            instance: {
              columns: [
                {
                  key: 'node',
                  name: t('节点名')
                }
              ],
              list: computerState.computer.list.data.records.map(d => ({
                node: d.metadata.name,
                isChecked:
                  !computerState.computer.selections.length ||
                  computerState.computer.selections.find(ins => ins.metadata.name === d.metadata.name)
              }))
            }
          }
        : {
            tables: [
              {
                table: 'k8s_pod',
                conditions: [
                  ['tke_cluster_instance_id', '=', this.props.route.queries.clusterId],
                  [
                    'node',
                    '=',
                    nodeName ||
                      (computerState.computer.list.data.records[0]
                        ? computerState.computer.list.data.records[0].metadata.name
                        : '')
                  ]
                ],
                fields: podMonitorFields
              }
            ],
            groupBy: [{ value: 'pod_name' }],
            instance: {
              columns: [
                {
                  key: 'pod_name',
                  name: t('Pod名称')
                }
              ],
              list: []
            }
          };

    this.setState({
      monitorPanelProps: {
        ...monitorPanelProps,
        title: t('节点监控'),
        headerExtraDOM: (
          <ul className="form-list">
            <li>
              <div className="form-label">
                <label>{t('对比维度')}</label>
              </div>
              <div className="form-input">
                <div className="form-unit">
                  <div className="tc-15-rich-radio" role="radiogroup">
                    {[
                      { label: t('节点'), key: 'nodeMonitor' },
                      { label: 'Pod', key: 'podMonitor' }
                    ].map(item => (
                      <button
                        key={item.key}
                        onClick={e => this._handleMonitor(item.key)}
                        className={'tc-15-btn m ' + (monitorType === item.key ? 'checked' : '')}
                        role="radio"
                      >
                        {item.label}
                      </button>
                    ))}
                  </div>
                </div>
              </div>
            </li>
            {monitorType === 'podMonitor' && (
              <li>
                <div className="form-label">
                  <label>{t('所属节点')}</label>
                </div>
                <div className="form-input">
                  <div className="form-unit">
                    <select
                      className="tc-15-select m"
                      onChange={e => {
                        this._handleMonitor(monitorType, e.target.value);
                      }}
                    >
                      {computerState.computer.list.data.records.map(item => (
                        <option key={item.metadata.name} value={item.metadata.name}>{`${item.metadata.name}`}</option>
                      ))}
                    </select>
                  </div>
                </div>
              </li>
            )}
          </ul>
        )
      }
    });
  }
}
