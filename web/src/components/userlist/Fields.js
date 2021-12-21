var dateFormat = (value) => {
    let d = new Date(value * 1000)
    return d.toLocaleString("ru")
}

const Fields = [{
        "name": "Username",
        "title": "Username",
        "width": "20%",
    },
    {
        "name": "Role",
        "title": "Role",
        "width": "10%",
    },
    {
        "name": "Email",
        "title": "Email",
        "visible": false,
    },
    {
        "name": "Register",
        "title": "Create date",
        "width": "50%",
        "formatter": dateFormat
    },
    {
        "name": "UUID",
        "title": "UUID",
        "visible": false,
    },
]

export default Fields