import { NavBar } from "./navbar"
import { useState} from "react"
export function Reporte(){
    const reporte = async(estructura) =>{
        const resp = await fetch("http://127.0.0.1:3000/Admin/Reporte",{
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify({
                "estructura":estructura
            })
        })
    }

    
    return (
        <>
        <NavBar/>
        <div className="container">
            <div>
                <button className="ebtn" type="button" onClick={ ()=>{reporte("ArbolB")}}>Arbol B</button>
                <button className="ebtn" type="button" onClick={() =>{reporte("Arbol-Merkle")}}> Arbol Merkle</button>
                <button className="ebtn" type="button" onClick={() =>{reporte("Grafo")}}>Grafo</button>
            </div>
        </div>
        </>
    )
}