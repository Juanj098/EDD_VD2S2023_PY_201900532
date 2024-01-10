import { NavBar } from "./navbar"
import { useState, useEffect } from "react"
import decline from "../assets/decline_icon.png"
import check from "../assets/accepted_icon.png"

export function ControlBooks(){
    const [Book, setBooks] = useState(null)
    useEffect(()=>{
        fetch("http://127.0.0.1:3000/Admin/LibrosPendientes")
        .then((response) => response.json())
        .then((data) => setBooks(data.LibrosPendientes))
    },[])
    console.log(Book)

    const stateBook = async(name,carnet,estado)=>{
        const resp = await fetch("http://127.0.0.1:3000/Admin/bookState",
        {
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body: JSON.stringify({
                "Name":name,
                "Estado":estado,
                "Tutor":parseInt(carnet)
            })
        })
    }

    const Finalizar = async() => {
        const resp = await fetch("http://127.0.0.1:3000/Admin/EndBooks",
        {
            method:"GET",
            headers:{
                "Content-Type":"application/json"
            }
        })
        console.log(resp)
    }

    return(
        <>
        <NavBar/>
        <div className="container">
            <table className="TBooksP">
                <thead>
                    <tr>
                        <th key="titulo-header">Titulo</th>
                        <th key="Tutor-header">Tutor</th>
                        <th key="aceptado-header"></th>
                        <th key="rechazado-header"></th>
                    </tr>
                </thead>
                <tbody>
                    {Book?.map((book)=>(
                        <tr key={book.Name}>
                            <td>{book.Name}</td>
                            <td>{book.Tutor}</td>
                            <td><button className="BtnCH" type="button" onClick={() => stateBook(book.Name,book.Tutor,"Aceptado")}><img className="checkbtn" src={check} alt="" /></button></td>
                            <td><button className="BtnDl" type="button" onClick={() => stateBook(book.Name,book.Tutor,"Rechazado")}><img className="declinebtn" src={decline} alt="" /></button></td>
                        </tr>
                    ))}
                </tbody>
            </table>
            <button type="button" className="endBtn">FINALIZAR</button>
        </div>
        </>
    )
}