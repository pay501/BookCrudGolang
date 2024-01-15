"use client"
import axios from "axios"
import { useEffect, useState } from "react"

export default function GetBook (){
    
    const [items, setItems] = useState([])

    useEffect(()=>{
        axios.get("http://localhost:8080/books")
        .then((res)=>{
            setItems(res.data)
        })
    },[])
    
    return(
        <div>
            {
                items.map((val, key)=>{
                    return(
                        <div key={key}>
                            {val.name}
                        </div>
                    )
                })
            }
        </div>
    )
}