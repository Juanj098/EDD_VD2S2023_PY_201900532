import { NavBar } from "./navbar"
import { useState } from "react"
export const  ReadCsv = () =>{
    const [tutores,setTutores] = useState(null)
    const [estudiantes,setEstudiantes] = useState(null)
    const [cursos, setCursos] = useState(null)
    
    const subirTutores =(e)=>{
        setTutores(e.target.files)
    }

    const subirEstudiantes = (e) => {
        setEstudiantes(e.target.files)
    }

    const subirCursos = (e) => {
        setCursos(e.target.files)
    }

    const PeticionTutores = async() => {
        const f = new FormData()
        for (const tutor of tutores){
            f.append('upload',tutor)
        }
        for (const pair of f.entries()) {
            console.log(pair[0], pair[1]);
        }
        try{
            const respT = await fetch("http://127.0.0.1:3000/Admin/TutoresCsv",
                {
                    method:'POST',
                    body:f
                })
                console.log(respT)
        }catch(error){
                
        }

    }
    
    const PeticionEstudiantes = async() => {
        const f = new FormData()
        for (const estudiante of estudiantes){
            f.append('upload',estudiante)
        }
        for (const pair of f.entries()) {
            console.log(pair[0], pair[1]);
        }
        try {
            const respE = await fetch("http://127.0.0.1:3000/Admin/EstudiantesCsv",
            {
                method:'POST',
                body: f
            })
            console.log(respE)
        } catch(error) {
            console.log(error)
        }
    }

    const PeticionCursos = async() => {
        const f = new FormData()
        for (const curso of cursos){
            f.append('upload',curso)
        }
        for (const pair of f.entries()){
            console.log(pair[0], pair[1]);
        }
        try {
            const respJson = await fetch("http://127.0.0.1:3000/Admin/CursosJSON",
            {
                method:'Post',
                body:f
            })
            console.log(respJson)
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <>
        <NavBar/>
        <div className="container">
            <h1>csv tutores y estudiantes...</h1>
            <form action="">
                <div className="csvT">
                    <label htmlFor="fileT"> TUTORES </label>
                    <input type="file" id='fileT' name="fileT" multiple  onChange={subirTutores}/>
                    <button className="csvtb" type="button" onClick={PeticionTutores}> cargar...</button>
                </div>
                <div className="csvE">
                    <label htmlFor="fileE"> ESTUDIANTES</label>
                    <input type="file" id='fileE' name="fileE" multiple onChange={subirEstudiantes}/>
                    <button className='csveb' type="button" onClick={PeticionEstudiantes} > cargar... </button>
                </div>
                <div className="cJSON">
                    <label htmlFor="jsonC">CURSOS</label>
                    <input type="file" name="jsonC" multiple onChange={subirCursos} />
                    <button className="jsonb" type="button" onClick={PeticionCursos}>cargar...</button>
                </div>
            </form>
        </div>
        </>
    )
} 