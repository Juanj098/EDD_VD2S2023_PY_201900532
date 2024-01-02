import { Link } from "react-router-dom"
import { useState } from "react"

export function Login(){
    const [data, setData] = useState({
        userlog:"",
        pass:"",
    })  

    let handlerLog = (e) =>{
        const {name,value} = e.target
        setData((prev)=>{
            return {...prev,[name]:value}
        })
    }  

    let handlerSubmit = async(e)=>{
        e.preventDefault()
        if (data.userlog.trim()!=="" && data.pass.trim()!=="" ){

            let obj = JSON.stringify(data)
            console.log(obj)
            try{
                const resp = await fetch("http://127.0.0.1:3000/Login",{
                    method:"POST",
                    headers:{
                        "Content-Type":"application/json",
                    },
                    body:obj,
                })
                const r = await resp.json()
                console.log(r.response)
                if (r.response == 'Admin'){
                    window.open("/Admin","_self")
                }else if (r.response == "Tutor"){
                    window.open("/Tutor","_self")
                } else if(r.response == 'Estudiante'){
                    window.open('/Estudiante',"_self")
                }else{
                    alert('verifique credenciales')
                }
            } catch(error){
                console.error(error.message)
            }
        } else{
            alert("Ingrese campos faltantes");
        }
    }

    return(
        <form className='bdy-log' onSubmit={handlerSubmit} >
            <div className='log-go'>
                <label className="lbl-log" htmlFor="userlog">User</label>
                <input className="input-txt" type="text" placeholder='Nombre de usuario' name="userlog" onChange={handlerLog} autoComplete="username" />
                <label className="lbl-log" htmlFor="pass">Password</label>
                <input className="input-txt" type="password" placeholder='Password' name="pass" onChange={handlerLog} autoComplete="current-password" />
            </div>
            <label className='a-btn-log'>
                <button  type="submit" className='btn-log'>Entrar</button>
            </label>
        </form>
    )}