import { YourRequestsContent } from '../../_pages/request'
import { connect } from 'react-redux'
import { requestActions } from '../../_actions'

const mapStateToProps = state => ({
    yourRequestRecords: state.request.yourRequestRecords,
    fetchStatus: state.request.fetchStatus
})

// 注入到展示组件的props中的回调方法
const mapDispatchToProps = dispatch => ({
    queryRequestsByUserIdAsync: (userId) => dispatch(requestActions.queryRequestsByUserIdAsync(userId)),
})

// 连接到展示组件
export const YourRequestsContentContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(YourRequestsContent)
