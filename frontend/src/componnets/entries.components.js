import React from 'react'
import {useState, useEffect} from 'react'
import axios from 'axios'
import {Button, Form, Container, Modal} from 'react-bootstrap'
import Entry from './single-entry.component'

const Entries = () => {
    const [entries, setEntries] = useState([])
    const [refreshData, setRefreshData] = useState(false)
    const [changeEntry, setChangeEntry] = useState({change: false, id: "0"})
    const [changeIngredient, setChangeIngredient] = useState({change: false, id: "0"})
    const [newIngredientName, setNewIngredientName] = useState("")
    const [addNewEntry , setAddNewEntry] = useState(false)
    const [newEntry, setNewEntry] = useState({
        "ingredients": "",
        "calories": 0,
        "fat": 0,
        "dish": ""
    })

    useEffect(() => {
        getAllEntries() 
    }, )
     
    if(refreshData){
        setRefreshData(false)
        getAllEntries()
    }

  return (
    <div>
        <Container><Button onClick={()=> setAddNewEntry(true)}  styles={{width: "100%" }}>Track Todays Calories</Button></Container>
        <Container>
            {entries != null && entries.map((entry, i) => <Entry key={i} deleteSingleEntry={deleteSingleEntry} setChangeIngredient={setChangeIngredient} setChangeEntry={setChangeEntry} entryData={entry} />)}
        </Container>

        <Modal show={addNewEntry} onHide={() => setAddNewEntry(false)} centered>
            <Modal.Header closeButton>
                <Modal.Title>Add New Entry</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Dish</Form.Label>
                        <Form.Control type="text" placeholder="Enter Dish" onChange={(e) => setNewEntry({...newEntry, "dish": e.target.value})} />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Ingredients</Form.Label>
                        <Form.Control type="text" placeholder="Enter Ingredients" onChange={(e) => setNewEntry({...newEntry, "ingredients": e.target.value})} />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Calories</Form.Label>
                        <Form.Control type="text" placeholder="Enter Calories" onChange={(e) => setNewEntry({...newEntry, "calories": e.target.value})} />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Fat</Form.Label>
                        <Form.Control type="number" placeholder="Enter Fat" onChange={(e) => setNewEntry({...newEntry, "fat": e.target.value})} />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => setAddNewEntry(false)}>
                    Close
                </Button>
                <Button variant="primary" onClick={addSingleEntry}>
                    Save Changes
                </Button>   
            </Modal.Footer>
        </Modal>

        <Modal show={changeIngredient.change} onHide={() => setChangeIngredient({change: false, id: "0"})}  centered>
            <Modal.Header closeButton>
                <Modal.Title>Change Ingredients</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Ingredients</Form.Label>
                        <Form.Control type="text" placeholder="Enter Ingredients" onChange={(e) => setNewIngredientName(e.target.value)} />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => setChangeIngredient({change: false, id: "0"})}>
                    Close
                </Button>
                <Button variant="primary" onClick={() => changeIngredientForEntry()}>
                    Save Changes
                </Button>   
            </Modal.Footer>
        </Modal>

        <Modal show={changeEntry.change} onHide={() => setChangeEntry({change: false, id: "0"})} centered>
            <Modal.Header closeButton>
                <Modal.Title>Change Entry</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Dish</Form.Label>
                        <Form.Control type="text" placeholder="Enter Dish" onChange={(e) => setNewEntry({...newEntry, "dish": e.target.value})} />
                        <Form.Label>Ingredients</Form.Label>
                        <Form.Control type="text" placeholder="Enter Dish" onChange={(e) => setNewEntry({...newEntry, "ingredients": e.target.value})} />
                        <Form.Label>Calories</Form.Label>
                        <Form.Control type="text" placeholder="Enter Dish" onChange={(e) => setNewEntry({...newEntry, "calories": e.target.value})} />
                        <Form.Label>Fat</Form.Label>
                        <Form.Control type="number" placeholder="Enter Dish" onChange={(e) => setNewEntry({...newEntry, "fat": e.target.value})} />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => setChangeEntry({change: false, id: "0"})}>
                    Close
                </Button>
                <Button variant="primary" onClick={() => changeSingleEntry()}>
                    Save Changes
                </Button>   
            </Modal.Footer> 
        </Modal>
    </div>
  )

  function changeSingleEntry() {
    changeEntry.change = false
    // setChangeEntry(changeEntry)
    var url= "http://localhost:8000/entry/update/"+changeEntry.id
    axios.put(url, newEntry)
    .then(Response => {
        if(Response.satus === 200){
            setRefreshData(true)
        }
    })
    .catch(error => {
        console.log(error)
    })
  }

  function changeIngredientForEntry() {
    changeIngredient.change = false
    var url= "http://localhost:8000/ingredient/update/"+changeIngredient.id
    axios.put(url, 
        {
            "ingredient": newIngredientName
        }
    )
    .then(Response => {
        if(Response.satus === 200){
            setRefreshData(true)
        }
    })
    .catch(error => {
        console.log(error)
    })
  }
  function addSingleEntry() {
    setAddNewEntry(false) 
    var url= "http://localhost:8000/entry/create"
    axios.post(url, {
        "ingredients": newEntry.ingredients,
        "calories": newEntry.calories,
        "fat":parseFloat( newEntry.fat),
        "dish": newEntry.dish
    })
    .then(Response => {
        if(Response.satus === 200){
            setRefreshData(true)
        }
    })
    .catch(error => {
        console.log(error)
    })
}

function deleteSingleEntry(id) {
    var url= "http://localhost:8000/entry/delete/"+id
    axios.delete(url, {
        
    }).then(Response=>{
        if(Response.status=== 200){
            setRefreshData(true)
        }
    })
}

function getAllEntries(){
    var url= "http://localhost:8000/entries"
    axios.get(url,{
        responseType: 'json'
    }).then(Response=>{
        if(Response.status===200){
            setEntries(Response.data)
        }
    })
}
}



export default Entries