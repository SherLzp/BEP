import { AllRequestsContent } from '../../_pages/request'
import { connect } from 'react-redux'
import { requestActions } from '../../_actions'

const mapStateToProps = state => ({
    allRequestRecords: state.request.allRequestRecords,
    fetchStatus: state.request.fetchStatus
})

// 注入到展示组件的props中的回调方法
const mapDispatchToProps = dispatch => ({
    showAllRequestsAsync: () => dispatch(requestActions.showAllRequestsAsync()),
})

// 连接到展示组件
export const AllRequestsContentContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(AllRequestsContent)
