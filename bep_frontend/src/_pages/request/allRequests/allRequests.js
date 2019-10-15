import React, { useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper'
import Container from '@material-ui/core/Container'
import { Button, Grid, DialogActions } from '@material-ui/core'
import CloudUploadIcon from '@material-ui/icons/CloudUpload'
import { FETCH_STATUS } from '../../../_constants'
import CircularProgress from '@material-ui/core/CircularProgress'
import TextField from '@material-ui/core/TextField'
import Dialog from '@material-ui/core/Dialog'
import DialogContent from '@material-ui/core/DialogContent'
import DialogTitle from '@material-ui/core/DialogTitle'
import Draggable from 'react-draggable'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Typography from '@material-ui/core/Typography'
import Switch from '@material-ui/core/Switch'
import { Link } from 'react-router-dom'

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
    input: {
        display: 'none',
    },
}))

function PaperComponent(props) {
    return (
        <Draggable cancel={'[class*="MuiDialogContent-root"]'}>
            <Paper {...props} />
        </Draggable>
    )
}

var description
var requestId
var answer

export function AllRequestsContent(props) {
    const classes = useStyles()
    const [open, setOpen] = React.useState(false)
    const [state, setState] = React.useState({
        checkedA: true,
        answer: "",
    })

    const handleChangeSwitch = name => event => {
        setState({ ...state, [name]: event.target.checked })
    }

    const handleChange = name => event => {
        answer = event.target.value
    }

    let rows = []
    useEffect(() => {
        props.showAllRequestsAsync()
        return () => {
        };
    }, [])

    if (props.allRequestRecords !== undefined) {
        console.log('allRequestRecords: ', props.allRequestRecords)
        rows = props.allRequestRecords
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

    const handleClickOpen = (row) => {
        setOpen(true)
        description = row.requirement
        requestId = row.request_id
    }

    const handleClose = () => {
        setOpen(false)
    }

    const onClickPushResponse = () => {
        console.log('answer: ', answer)
        props.pushResponseAsync(requestId, 'Jack', answer)
        if (props.fetchStatus == FETCH_STATUS.FETCH_SUCCESS) {
            alert('push response success')
        }
    }

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Paper className={classes.root}>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>Description</TableCell>
                            <TableCell align="right">Owner</TableCell>
                            <TableCell align="right">Reward</TableCell>
                            <TableCell align="right">Created Time</TableCell>
                            <TableCell align="right">Expired Time</TableCell>
                            <TableCell align="right">Status</TableCell>
                            <TableCell align="right">Answer it</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map(row => (
                            <TableRow key={row.name}>
                                <TableCell component="th" scope="row">
                                    {row.requirement}
                                </TableCell>
                                <TableCell align="right">{row.user_id}</TableCell>
                                <TableCell align="right">{row.reward} $</TableCell>
                                <TableCell align="right">{row.create_time}</TableCell>
                                <TableCell align="right">{row.expired_time}</TableCell>
                                <TableCell align="right">{row.status === 0 ? "Not Completed" : row.status === 1 ? "Completed" : "Expired"}</TableCell>
                                <TableCell align="right">
                                    {row.status === 0 ? <Button color="primary" onClick={() => handleClickOpen && handleClickOpen(row)}>
                                        Response
                                        <CloudUploadIcon className={classes.rightIcon} />
                                    </Button>
                                        :
                                        <Button color="primary" disabled>
                                            Response
                                        <CloudUploadIcon className={classes.rightIcon} />
                                        </Button>
                                    }

                                </TableCell>
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
                        <Typography align="center" variant="h4">Push Response</Typography>
                    </DialogTitle>
                    <DialogContent>
                        <Grid>
                            <TextField
                                id="standard-full-width"
                                label="Request Description"
                                style={{ margin: 8 }}
                                value={description}
                                fullWidth
                                margin="normal"
                                InputLabelProps={{
                                    shrink: true,
                                }}
                            />
                        </Grid>
                        <Grid>
                            <TextField
                                id="standard-full-width"
                                label="Your Response"
                                style={{ margin: 8 }}
                                placeholder="Just put your response"
                                fullWidth
                                margin="normal"
                                onChange={handleChange('answer')}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                            />
                        </Grid>
                        {/* <label>or</label>
                        <Grid>
                            <input
                                accept="file/*"
                                className={classes.input}
                                id="outlined-button-file"
                                multiple
                                type="file"
                            />
                            <label htmlFor="outlined-button-file">
                                <Button variant="outlined" component="span" className={classes.button}>
                                    Upload
                                </Button>
                            </label>
                        </Grid> */}
                        <Grid>
                            <FormControlLabel
                                control={
                                    <Switch checked={state.checkedA} onChange={handleChangeSwitch('checkedA')} value="checkedA" />
                                }
                                label="Keep anonymous"
                            />
                        </Grid>
                    </DialogContent>
                    <DialogActions>
                        <Button variant="contained" color="primary" onClick={onClickPushResponse}>Push Response</Button>
                    </DialogActions>
                </Dialog>
            </Paper>
        </Container>
    )
}