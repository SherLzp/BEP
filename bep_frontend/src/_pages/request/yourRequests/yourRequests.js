import React, { useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import Container from '@material-ui/core/Container'
import { Button } from '@material-ui/core'
import { Link } from 'react-router-dom'
import { FETCH_STATUS } from '../../../_constants'
import CircularProgress from '@material-ui/core/CircularProgress'

const useStyles = makeStyles(theme => ({
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
}))

function createData(description, reward, status, expired, responsesCount, acceptResponseId, requestId) {
    return { description, reward, status, expired, responsesCount, acceptResponseId, requestId }
}

// const rows = [
//     createData(`Is anyone has Lin's Passport?`, 10.0, 'Not completed', '2019-09-01', 3, 'Response_001', '001'),
//     createData(`Is anyone has Lin's Driver License?`, 50.0, 'Failed', '2019-08-01', 0, '-', '002'),
//     createData(`Is anyone has He's Passport?`, 30.0, 'Not completed', '2019-09-02', 1, '-', '003'),
// ]

export function YourRequestsContent(props) {
    const classes = useStyles();

    let rows = []
    useEffect(() => {
        props.queryRequestsByUserIdAsync('Sher')
        return () => {
        };
    }, [])

    if (props.yourRequestRecords !== undefined) {
        console.log('yourRequestRecords: ', props.yourRequestRecords)
        rows = props.yourRequestRecords
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

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid>
                <Button color="primary" variant="contained" component={Link} to="/requests/pushRequest">Create Request</Button>
            </Grid>
            <br></br>
            <Grid>
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
                                <TableCell align="right">Responses</TableCell>
                                <TableCell align="right">AcceptResponse</TableCell>
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
                                    <TableCell align="right">{row.responses.length}</TableCell>
                                    <TableCell align="right">{row.accept_response_id}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
            </Grid>
        </Container>
    );
}