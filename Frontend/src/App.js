import './App.css';
import * as React from 'react';
import { useState } from "react";
import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import axios from 'axios';
import Box from '@mui/material/Box';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import { styled } from '@mui/material/styles';
import { DataGrid } from '@mui/x-data-grid';

const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'left',
    color: theme.palette.text.secondary,
}));

function App() {
    const [rows, setRows] = useState([])
    const [showText, setShowText] = useState("")
    const [fields, setValues] = useState({
        "business_name": "",
        "established_year": '',
        "loan_amount": "",
        "account_provider": "",
    });
    const handleFieldChange = event => {
        setValues({
            ...fields,
            [event.target.id]: event.target.value
        });
    }
    const onSubmitBs = () => {
        setShowText("")
        const emptyKeys = Object.keys(fields).reduce((next, field) => {
            if (fields[field] && fields[field].trim() == "") {
                return { ...next, [field]: "" };
            }
        }, {});


        setValues({
            ...fields,
            ...emptyKeys,
        })
        const config = {
            headers: {
                'content-type': 'multipart/form-data',
                'X-API-KEY': 'super-secret'
            }
        }

        axios.post(`http://localhost:8080/api/v1/${fields.business_name}/balancesheet/${fields.account_provider}`, fields, config)
            .then(response => {
                if (response && response.data && response.data.length > 0) {
                    setRows(response.data)
                    return
                }
                setRows([])
                setShowText("No balance sheet found for specified business and Account provider")
            })
            .catch(error => {
                console.log(error);
                setRows([])
                setShowText(error.message)
            });
    }

    const onSubmit = () => {
        setShowText("")
        const emptyKeys = Object.keys(fields).reduce((next, field) => {
            if (fields[field] && fields[field].trim() == "") {
                return { ...next, [field]: "" };
            }
        }, {});


        setValues({
            ...fields,
            ...emptyKeys,
        })
        const config = {
            headers: {
                'content-type': 'multipart/form-data',
                'X-API-KEY': 'super-secret'
            }
        }

        axios.post(`http://localhost:8080/api/v1/${fields.business_name}/submit`, fields, config)
            .then(response => {
                let txt = "Your loan is not approved"
                if (response && response.data && response.data.decision) {
                    txt = "Your loan is approved for amount " + response.data.approved_amount
                }
                setShowText(txt)
            })
            .catch(error => {
                console.log(error);
                setShowText(error.message)
            });
    }

    return (
        <div >
            <div style={{ flexGrow: 1, display: "grid", justifyContent: "center" }} >
                <h2>Loan Application</h2>
                <h3 style={{ background: "yellow", display: showText != "" ? "block" : "none" }}>{showText}</h3>
                <Grid container spacing={2} style={{ width: "850px" }}>
                    <Grid item xs={8} style={{ justifyContent: "center" }}>
                        <Item>
                            <TextField
                                required
                                error={fields.name == "" ? true : false}
                                id="business_name"
                                label="Business Name"
                                value={fields.business_name}
                                onChange={handleFieldChange}
                            />
                            <br />
                            <br />
                            <TextField
                                type="number"
                                required
                                error={fields.email == "" ? true : false}
                                id="established_year"
                                label="Established Year"
                                value={fields.established_year}
                                onChange={handleFieldChange}
                            />
                            <br />
                            <br />
                            <TextField
                                required
                                type="number"
                                error={fields.phone == "" ? true : false}
                                id="loan_amount"
                                label="Loan Amount"
                                value={fields.loan_amount}
                                onChange={handleFieldChange}
                            />
                            <br />
                            <br />
                            <FormControl fullWidth>
                                <InputLabel id="demo-simple-select-label">Select Account Provider</InputLabel>
                                <Select
                                    labelId="account_provider"
                                    id="account_provider"
                                    value={fields.account_provider}
                                    label="Select Account Provider"
                                    onChange={(e) => {
                                        e.target.id = "account_provider"
                                        handleFieldChange(e)
                                    }}
                                >
                                    <MenuItem value="XERO">Xero</MenuItem>
                                    <MenuItem value="MYOB">MYOB</MenuItem>
                                </Select>
                            </FormControl>
                        </Item>
                        <br />
                        <Button variant="contained" onClick={onSubmitBs} style={{ float: "right" }}  >Balance Sheet</Button> &nbsp;
                    </Grid>
                </Grid>
            </div>
            <div style={{ margin: "20px" }}>
                <div>
                    {rows && rows.length > 0 ? <Button variant="contained" style={{ float: "right", marginRight: "10px", paddingBottom: "10px" }} onClick={onSubmit}>Submit Application</Button> : <></>}
                </div>
                <div>
                    <h3>Review Balance sheet and Summit loan application:</h3>
                    <DataGrid
                        rows={rows}
                        columns={columns}
                        initialState={{
                            pagination: {
                                paginationModel: { page: 0, pageSize: 100 },
                            },
                        }}
                        pageSizeOptions={[5, 100]}
                        getRowId={(row) => row.year + row.month}
                    />
                </div>
            </div>
        </div>
    );
}

const columns = [
    { field: 'year', headerName: 'Year', type: 'number', width: 130 },
    { field: 'month', headerName: 'Month', type: 'number', width: 130 },
    { field: 'profitOrLoss', headerName: 'Profit Or Loss', type: 'number', width: 130 },
    { field: 'assetsValue', headerName: 'Assets Value', type: 'number', width: 130 },
];

export default App;
