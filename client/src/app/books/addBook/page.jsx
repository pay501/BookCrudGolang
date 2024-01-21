"use client"
import axios from "axios";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import Swal from "sweetalert2";
import style from "./addBook.module.css";

export default function AddBookPage() {
    const router = useRouter()
    const [form, setForm] = useState({
        name: "",
        author: "",
        description: "",
        price: 0.0,
        image: "",
    })

    async function checkSession() {
        try {
            await axios.get("http://localhost:8080/book/1", { withCredentials: true })
                .then(res => {
                    if (res.data.message === "return") {
                        router.push("/login")
                    }
                })
        } catch (err) {
            router.push("/login")
        }
    }

    function inputChange(e) {
        e.preventDefault();
        const { name, value } = e.target
        setForm((prevForm) => ({ ...prevForm, [name]: value }))
    }

    function addBook(e) {
        e.preventDefault();
        if (form.name === "" || form.author === "" || form.description === "" || form.price == 0.0 || form.image === "") {
            Swal.fire({
                icon: "error",
                title: "Oops! Can't add book!",
                text: "Please ,fill in all data"
            })
        } else {
            axios.post("http://localhost:8080/addBook", {
                name: form.name,
                author: form.author,
                description: form.description,
                price: parseFloat(form.price),
                image: form.image
            }, { withCredentials: true })
                .then(res => {
                    if (res.data.message === "Add successful") {
                        Swal.fire({
                            icon: "success",
                            title: "Book has added already!",
                            footer: "Thank you!"
                        })
                        router.push("/")
                    } else {
                        Swal.fire({
                            icon: "error",
                            title: "There is something wrong, can't add book.",
                            footer: "Thank you!"
                        })
                    }
                })
        }

    }

    useEffect(() => {
        checkSession()
    }, [])
    return (
        <div>
            <div className={style.container}>
                <div className={`${style.imgContainer}`}>
                    <Image
                        src="https://images.pexels.com/photos/1907785/pexels-photo-1907785.jpeg?auto=compress&cs=tinysrgb&w=600"
                        className={style.img}
                        alt=""
                        fill
                    >
                    </Image>
                </div>
                <div className={`${style.formContainer}`}>
                    <div>
                        <h1 className="text-4xl font-bold">Welcome to the Library!</h1>
                        <p
                            className="text-lg font-semibold mt-4"
                        >
                            I am so appreciate that you were join our library.
                            <br /><span>
                                Wanna add your book?
                            </span>
                        </p>

                    </div>
                    <div>
                        <form className="flex flex-col gap-4 my-8">
                            <div>
                                <input
                                    type="text"
                                    placeholder='Title'
                                    name='name'
                                    className="input input-primary w-full max-w-md"
                                    onChange={inputChange}
                                    required
                                />
                            </div>
                            <div>
                                <input
                                    type="text"
                                    placeholder='Author'
                                    name='author'
                                    className="input input-primary w-full max-w-md"
                                    onChange={inputChange}
                                    required
                                />
                            </div>
                            <div>
                                <input
                                    type="text"
                                    placeholder='Description'
                                    name='description'
                                    className="input input-primary w-full max-w-md"
                                    onChange={inputChange}
                                    required
                                />
                            </div>
                            <div>
                                <input
                                    type="number"
                                    placeholder='Price'
                                    name='price'
                                    className="input input-primary w-full max-w-md"
                                    onChange={inputChange}
                                    required
                                />
                            </div>
                            <div>
                                <input
                                    type="text"
                                    placeholder='Image URL'
                                    name='image'
                                    className="input input-primary w-full max-w-md"
                                    onChange={inputChange}
                                    required
                                /><br />
                                <label className="text-sm text-red-500">*Sorry for inconvenience*</label>
                            </div>
                            <button
                                className={`bg-blue-500 rounded-lg py-1 w-full max-w-md text-white font-semibold h-10`}
                                onClick={addBook}
                            >
                                Add
                            </button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}