import { useEffect, useState } from "react"

export function useFetch(url){
    const [Data, setData] = useState(null)
    useEffect(()=>{
        fetch(url)
        .then((response)=>response.json())
        .then((Data)=>setData(Data))
    },[])
    return {Data}
}


