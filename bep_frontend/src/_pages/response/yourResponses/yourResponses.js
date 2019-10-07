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

function createData(requestDesc, requestOwner, status, expired, response, requestId) {
    return { requestDesc, requestOwner, status, expired, response, requestId }
}

// const rows = [
//     createData(`Is anyone has He's Passport?`, 'UnKnown', 'Not completed', '2019-09-02', 'http://img.sher.vip/1.jpg', '003'),
//     createData(`Is anyone has He's Driver License?`, 'UnKnown', 'Not completed', '2019-09-03', 'http://img.sher.vip/1.jpg', '004'),
//     createData(`How to go to Hangzhou?`, 'Sher', 'Over', '2019-09-11', 'Just take railway', '005'),
// ]

export function YourResponsesContent(props) {
    const classes = useStyles()

    let rows = []
    useEffect(() => {
        props.queryResponsesByUserIdAsync('Jack')
        return () => {
        };
    }, [])

    if (props.yourResponseRecords !== undefined) {
        console.log('yourResponseRecords: ', props.yourResponseRecords)
        rows = props.yourResponseRecords
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
            <Paper className={classes.root}>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>Request Description</TableCell>
                            <TableCell align="right">Request Owner</TableCell>
                            <TableCell align="right">Expire Date</TableCell>
                            <TableCell align="right">Status</TableCell>
                            <TableCell align="right">Response</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map(row => (
                            <TableRow key={row.name}>
                                <TableCell component="th" scope="row">
                                    {row.requestDesc}
                                </TableCell>
                                <TableCell align="right">{row.requestOwner}</TableCell>
                                <TableCell align="right">{row.expired}</TableCell>
                                <TableCell align="right">{row.status}</TableCell>
                                <TableCell align="right">
                                    11
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </Paper>
        </Container>
    )
}