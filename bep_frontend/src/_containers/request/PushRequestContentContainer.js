import { PushRequestContent } from '../../_pages/request'
import { connect } from 'react-redux'
import { requestActions } from '../../_actions'

const mapStateToProps = state => ({
    fetchStatus: state.request.fetchStatus
})

// 注入到展示组件的props中的回调方法
const mapDispatchToProps = dispatch => ({
    pushRequestAsync: (userId, requirement, reward, expiredTime) => dispatch(requestActions.pushRequestAsync(userId, requirement, reward, expiredTime)),
})

// 连接到展示组件
export const PushRequestContentContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(PushRequestContent)
