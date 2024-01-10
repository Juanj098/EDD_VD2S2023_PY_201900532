import { NavBarT } from "./navbarT"
import { useState } from "react"
export function PublicacionesT(){
    const [data, setData] = useState(null)

    const HandlesText = (e) => {
        setData(e.target.value)
    }

    const Publicar = async(e) =>{
        e.preventDefault()
        if (data != ''){
            try{
                let resp = await fetch("http://127.0.0.1:3000/Tutor/CargarPublicacion",
                {
                    method:'POST',
                    headers:{
                        "Content-Type":"application/json"
                    },
                    body: JSON.stringify({
                        "Tutor":parseInt(localStorage.getItem("user")),
                        "Publicacion": data
                    })
                })
                console.log(resp)
            } catch(error){
                console.log(error)
            }
        } else{
            alert('Ingrese Publicacion')
        }
    }

    return(
        <>
            <NavBarT/>
            <div className="container">
            <textarea className="txtPubli" rows="25" cols="100" placeholder="Publicacion..." onChange={HandlesText}></textarea>
            <button className="btnTxt" type="button" onClick={Publicar}>Publicar</button>
            </div>
        </>
    )
}