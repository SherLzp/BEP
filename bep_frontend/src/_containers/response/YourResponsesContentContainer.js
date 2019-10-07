import { YourResponsesContent } from '../../_pages/response'
import { connect } from 'react-redux'
import { responseActions } from '../../_actions'

const mapStateToProps = state => ({
    yourResponseRecords: state.response.yourResponseRecords,
    fetchStatus: state.request.fetchStatus
})

// 注入到展示组件的props中的回调方法
const mapDispatchToProps = dispatch => ({
    queryResponsesByUserIdAsync: (userId) => dispatch(responseActions.queryResponsesByUserIdAsync(userId)),
})

// 连接到展示组件
export const YourResponsesContentContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(YourResponsesContent)
