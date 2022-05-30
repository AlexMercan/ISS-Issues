import React, { useState } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import * as auth from '../services/AuthService'
import { useNavigate } from 'react-router';
import { FormControlLabel, FormLabel, Radio, RadioGroup } from '@mui/material';
import { Label } from '@mui/icons-material';

const theme = createTheme();

function validateCredentials(username: String, password: String) {
    let errors: String[] = []
    if (username.length <= 5 || username.length >= 13) {
        errors.push("Username length")
    }
    if (password.length < 8 || password.length > 15) {
        errors.push("Password length")
    }
    if (errors.length === 0) {
        return ""
    }
    return "Invalid " + errors.join(" and ");
}

export default function Register() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [role, setRole] = useState("Tester")
    const [validationError, setValidationError] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        let erorrStr = validateCredentials(username, password)
        if (erorrStr !== "") {
            setValidationError(erorrStr)
            return;
        }
        var res = await auth.Register(username, password, role)
        if (res == true) {
            navigate("/login");
        }
    };

    return (
        <ThemeProvider theme={theme}>
            <Container component="main" maxWidth="xs">
                <CssBaseline />
                <Box
                    sx={{
                        marginTop: 8,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    <Avatar sx={{ m: 1, bgColor: 'secondary.main' }}>
                        <LockOutlinedIcon />
                    </Avatar>
                    <Typography component="h1" variant="h5">
                        Register
                    </Typography>
                    <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <TextField
                                    required
                                    fullWidth
                                    id="username"
                                    label="Username"
                                    name="username"
                                    autoComplete="username"
                                    onChange={(e) => { setUsername(e.target.value) }}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField
                                    required
                                    fullWidth
                                    name="password"
                                    label="Password"
                                    type="password"
                                    id="password"
                                    autoComplete="new-password"
                                    onChange={(e) => { setPassword(e.target.value) }}
                                />
                            </Grid>
                        </Grid>
                        <FormLabel id="role-radio" sx={{ fontSize: 18 }}>Role</FormLabel>
                        <RadioGroup
                            aria-labelledby="role-radio"
                            name="role-radio-group"
                            value={role}
                            onChange={(e) => { setRole(e.target.value) }}
                        >
                            <FormControlLabel value="Tester" control={<Radio />} label="Tester" />
                            <FormControlLabel value="Programmer" control={<Radio />} label="Programmer" />
                        </RadioGroup>
                        {validationError !== "" && <span style={{ "color": "red" }}>{validationError}</span>}
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign Up
                        </Button>
                        <Grid container justifyContent="center">
                            <Grid item>
                                <Link href="/login" variant="body2">
                                    Already have an account? Sign in
                                </Link>
                            </Grid>
                        </Grid>
                    </Box>
                </Box>
            </Container>
        </ThemeProvider>
    );
}
