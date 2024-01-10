import { NavBarE } from "./navBarE"
import { useState , useEffect } from "react"
export function VisualizarLibros(){
    // const[Cursos, setCursos ]= useState(null)
    // const[Books, setBooks] = useState(null)
    
    // useEffect(()=>{
    //     fetch('http://127.0.0.1:3000/Estudiante/Cursos',{
    //         method:'POST',
    //         headers:{
    //             "Content-Type":"application/json"
    //         },
    //         body:JSON.stringify({
    //             "carnet":parseInt(localStorage.getItem('user'))
    //         })
    //     })
    //     .then((response) => response.json())
    //     .then((curso) => setCursos(curso.Arreglo))
    

    return (
        <>
            <NavBarE/>
            <div className="container">
                <h1>lIBROS DE TUTORES</h1>
                <div className="containerCards">
                    {/* {Cursos?.map((curso,index)=>(
                    <div key={index} className="card" >
                        <h1>{curso}</h1>
                    </div> */}
                    {/* ))} */}
                </div>
            </div>
        </>
    )
}