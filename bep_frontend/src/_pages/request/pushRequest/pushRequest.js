import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import TextField from '@material-ui/core/TextField'
import Grid from '@material-ui/core/Grid'
import Container from '@material-ui/core/Container'
import { Button, Typography } from '@material-ui/core'
import { Link } from 'react-router-dom'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Switch from '@material-ui/core/Switch'
import { FETCH_STATUS } from '../../../_constants'

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

export function PushRequestContent(props) {
    const classes = useStyles()
    const [values, setValues] = React.useState({
        checkedA: true,
        description: "",
        expiredTime: "2019-10-24T10:30",
        reward: 0,
    })

    const handleChange = name => event => {
        setValues({ ...values, [name]: event.target.value })
    }

    const handleChangeSwitch = name => event => {
        setValues({ ...values, [name]: event.target.checked })
    }

    const formatDate = (expiredTime) => {
        if (expiredTime !== undefined && expiredTime.length === 16) {
            return expiredTime.substr(0, 10) + ' ' + expiredTime.substr(11) + ':00'
        }
    }

    const onClickPushRequest = () => {
        console.log('expiredTime: ',values.expiredTime)
        let expiredTime = formatDate(values.expiredTime)
        props.pushRequestAsync('Sher', values.description, values.reward, expiredTime)
        if(props.fetchStatus == FETCH_STATUS.FETCH_SUCCESS){
            alert('push request success')
        }
    }

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid>
                <Typography variant="h4">Push Request</Typography>
            </Grid>
            <Grid>
                <TextField
                    id="standard-full-width"
                    label="Request Description"
                    style={{ margin: 8 }}
                    placeholder="Just like Is anyone has Lin's passport?"
                    fullWidth
                    margin="normal"
                    onChange={handleChange('description')}
                    InputLabelProps={{
                        shrink: true,
                    }}
                />
            </Grid>
            <Grid>
                <TextField
                    id="outlined-number"
                    label="reward"
                    onChange={handleChange('reward')}
                    type="number"
                    defaultValue={0}
                    className={classes.textField}
                    InputLabelProps={{
                        shrink: true,
                    }}
                    margin="normal"
                    variant="outlined"
                />
            </Grid>
            <Grid>
                <TextField
                    id="datetime-local"
                    label="Expired Time"
                    type="datetime-local"
                    defaultValue="2019-10-24T10:30"
                    className={classes.textField}
                    onChange={handleChange('expiredTime')}
                    InputLabelProps={{
                        shrink: true,
                    }}
                />
            </Grid>
            <Grid>
                <FormControlLabel
                    control={
                        <Switch checked={values.checkedA} onChange={handleChangeSwitch('checkedA')} value="checkedA" />
                    }
                    label="Keep anonymous"
                />
            </Grid>
            <Grid>
                <Button variant="contained" color="primary" onClick={onClickPushRequest}>Push Request</Button>
            </Grid>
        </Container >
    )
}