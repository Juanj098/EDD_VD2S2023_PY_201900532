
import { NavBarT } from "./navbarT"
import { useState } from "react"
export function LibrosT(){
    const [Pdf, SetPdf] = useState(null)
    const [PdfUrl, setPdfUrl]=useState(null)

    const handlerFile = (e) => {
        const file = e.target.files[0]
        const fileUrl = URL.createObjectURL(file)
        setPdfUrl(fileUrl)
        SetPdf(file)
    }

    const subirPdf = async () => {
        if (Pdf){
            console.log(Pdf)
            const respL = await fetch("http://127.0.0.1:3000/Tutor/CargarLibro",
            {
                method:"POST",
                headers:{
                    "Content-Type":"application/json"
                },
                body: JSON.stringify({
                    "Date":"",
                    "Name":Pdf.name,
                    "Estado":"Pendiente",
                    "Tutor":parseInt(localStorage.getItem("user"))
                })
            })
            console.log(respL)
        } else{
            console.log("no escogiste ningun libro :p")
        }
    }

    return(
        <>
            <NavBarT/>
            <div className="container">
                <div className="PdfT">
                    <label htmlFor="libro">Cargar libros</label>
                    <input id="libro" type="file" name="libro" multiple onChange={handlerFile} />
                    <button className="pdfb" type="button" onClick={subirPdf}>save</button>
                </div>
                <iframe className="pdfR" src={PdfUrl} frameBorder="0"></iframe>
            </div>
        </>
    )
}