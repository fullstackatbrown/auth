import React, { useEffect, useState } from 'react';
import './App.css';
import { Button, Table, TableBody, TableCell, TableHead, TableRow, TextField } from '@mui/material';

function App() {

  // Placeholder data
  let rows = [
    {course: 'S22CSCI0200', role: 'UTA'},
    {course: 'S22CSCI1951A', role: 'Student'},
    {course: 'S22CSCI1951I', role: 'Student'}
  ]

  const [currentEmail, setCurrentEmail] = useState('');
  const [currentID, setCurrentID] = useState('');
  const [roles, setRoles] = useState(rows);

  // API Call: GET /users to convert email to userID
  useEffect(() => {
    if (currentEmail !== '') {
      // TODO: Convert email to userID
      fetch(`http://localhost:8000/users/${currentEmail}`)
        .then(res => res.json())
        .then(data => {
          setCurrentID(data.id);
          // setCurrentID("googleIDplaceholder123");
        })
        .catch(err => console.log(err));
    }
  }, [currentEmail])

  // API Call: GET /users/{userId}/roles to get roles for a user
  useEffect(() => {
    if (currentID !== '') {
      fetch(`http://localhost:8000/users/${currentID}/roles`)
        .then(res => res.json())
        .then(data => {
          console.log("Getting roles for ID: " + currentID)
          console.log(data);
          // TODO: make sure that data is in the format [{course: string, role: string}]
          let fetchedRoles = [];
          for (let i = 0; i < data.length; i++) {
            fetchedRoles.push({course: data[i].domain, role: data[i].role});
          }
          setRoles(fetchedRoles);
        })
        .catch(err => console.log(err));
    }
  }, [currentID])

  // API Call: POST /users/{userId}/roles to add a role for a user
  const addRole = (userId: string, domain: string, role: string) => {
    // TODO: add role to database
  }

  // API Call: DELETE /users/{userId}/roles to delete a role for a user
  const deleteRole = (userId: string, domain: string, role: string) => {
    // TODO: delete role from database
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>Auth Roles</h1>
      </header>
      <div className="App-body">
        <TextField className='user-id-field' id="user-input" label="User ID" variant="outlined" value={currentEmail}
        onChange={(e) => setCurrentEmail(e.target.value)}/>
        {/* <Button variant="contained">Submit</Button> */}
        <Table sx={{ minWidth: 650, maxWidth: 1000 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Course</TableCell>
            <TableCell>Role</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {roles.map((row) => (
            <TableRow
              key={row.course}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {row.course}
              </TableCell>
              <TableCell>{row.role}</TableCell>
              <TableCell align='right'>
                <Button variant="contained" color="error" onClick={
                  () => deleteRole(currentID, row.course, row.role)
                }>Delete</Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <div className="add-role">
        <TextField className='course-field' id="course-input" label="Course" variant="outlined"/>
        <TextField className='role-field' id="role-input" label="Role" variant="outlined"/>
        <Button variant="contained" onClick={
          () => addRole(currentID, 'course', 'role')
        }>Add Role</Button>
      </div>
    </div>
    </div>
  );
}

export default App;
