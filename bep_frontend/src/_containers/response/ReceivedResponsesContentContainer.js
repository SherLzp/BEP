import { ReceivedResponsesContent } from '../../_pages/response'
import { connect } from 'react-redux'
import { responseActions } from '../../_actions'

const mapStateToProps = state => ({
    requestAndResponsesRecords: state.response.requestAndResponsesRecords,
    fetchStatus: state.response.fetchStatus
})

// 注入到展示组件的props中的回调方法
const mapDispatchToProps = dispatch => ({
    queryReceivedResponsesByUserIdAsync: (userId) => dispatch(responseActions.queryReceivedResponsesByUserIdAsync(userId)),
    acceptResponseAsync: (userId,requestId,responseId) => dispatch(responseActions.acceptResponseAsync(userId,requestId,responseId)),
})

// 连接到展示组件
export const ReceivedResponsesContentContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(ReceivedResponsesContent)
