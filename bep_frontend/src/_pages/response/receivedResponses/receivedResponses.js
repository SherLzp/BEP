import React, { useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper'
import Container from '@material-ui/core/Container'
import { FETCH_STATUS } from '../../../_constants'
import CircularProgress from '@material-ui/core/CircularProgress'
import { Button } from '@material-ui/core'
import Dialog from '@material-ui/core/Dialog'
import DialogContent from '@material-ui/core/DialogContent'
import DialogTitle from '@material-ui/core/DialogTitle'
import Draggable from 'react-draggable'
import Typography from '@material-ui/core/Typography'
import { PushRequestContent } from '../../request'

const useStyles = makeStyles(theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
    content: {
        flexGrow: 1,
        height: '100vh',
        overflow: 'auto',
    },
    container: {
        paddingTop: theme.spacing(4),
        paddingBottom: theme.spacing(4),
    },
    paper: {
        padding: theme.spacing(2),
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
    },
    fixedHeight: {
        height: 240,
    },
    rightIcon: {
        marginLeft: theme.spacing(1),
    },
}))

function PaperComponent(props) {
    return (
        <Draggable cancel={'[class*="MuiDialogContent-root"]'}>
            <Paper {...props} />
        </Draggable>
    )
}

var responses = []
var request = []
export function ReceivedResponsesContent(props) {
    const classes = useStyles()
    const [open, setOpen] = React.useState(false)
    const handleClickOpen = (row) => {
        setOpen(true)
        responses = row.responses
        request = row.request
    }

    const handleClose = () => {
        setOpen(false)
    }

    let rows = []
    useEffect(() => {
        props.queryReceivedResponsesByUserIdAsync('Sher')
        return () => {
        };
    }, [])

    if (props.requestAndResponsesRecords !== undefined) {
        console.log('all of your requests and responses: ', props.requestAndResponsesRecords)
        rows = props.requestAndResponsesRecords
    }

    if (props.fetchStatus == FETCH_STATUS.FETCH_BEGIN) {
        console.log("")
        return (
            <div align="center">
                <br />
                <CircularProgress />
            </div>
        )
    }

    const onClickAcceptResponse = (requestId, responseId) => {
        props.acceptResponseAsync('Sher', requestId, responseId)
        if (props.fetchStatus == FETCH_STATUS.FETCH_SUCCESS) {
            alert('accept response success')
        }
    }

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Paper className={classes.root}>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>Request Description</TableCell>
                            <TableCell align="right">Created Time</TableCell>
                            <TableCell align="right">Expired Time</TableCell>
                            <TableCell align="right">Responses Count</TableCell>
                            <TableCell align="right">Status</TableCell>
                            <TableCell align="right">Detail</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map(row => (
                            <TableRow key={row.name}>
                                <TableCell component="th" scope="row">
                                    {row.request.requirement}
                                </TableCell>
                                <TableCell align="right">{row.request.create_time}</TableCell>
                                <TableCell align="right">{row.request.expired_time}</TableCell>
                                <TableCell align="right">{row.responses === null ? 0 : row.responses.length}</TableCell>
                                <TableCell align="right">{row.request.status === 0 ? "Not Completed" : row.request.status === 1 ? "Completed" : "Expired"}</TableCell>
                                <TableCell align="right"><Button size="medium" color="primary" variant="contained" onClick={() => handleClickOpen && handleClickOpen(row)}>Detail Info</Button></TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
                <Dialog
                    open={open}
                    onClose={handleClose}
                    PaperComponent={PaperComponent}
                    aria-labelledby="draggable-dialog-title"
                    maxWidth={100}
                >
                    <DialogTitle id="draggable-dialog-title">
                        <Typography align="center" variant="h6">Responses</Typography>
                    </DialogTitle>
                    <DialogContent>
                        <Table className={classes.table}>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Response Answer</TableCell>
                                    <TableCell align="right">Created Time</TableCell>
                                    <TableCell align="right">User</TableCell>
                                    <TableCell align="right">Operation</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {responses === null ? "" : responses.map(row => (
                                    <TableRow key={row.name}>
                                        <TableCell component="th" scope="row">
                                            {row.answer}
                                        </TableCell>
                                        <TableCell align="right">
                                            {row.create_time}
                                        </TableCell>
                                        <TableCell align="right">
                                            {row.user_id}
                                        </TableCell>
                                        <TableCell align="right">
                                            {request.status === 0 ?
                                                <Button size="medium" color="primary" variant="contained" onClick={() => onClickAcceptResponse && onClickAcceptResponse(request.request_id, row.response_id)}>Accept</Button> :
                                                <Button size="medium" color="primary" variant="contained" disabled>Accept</Button>
                                            }
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </DialogContent>
                </Dialog>
            </Paper>
        </Container>
    )
}