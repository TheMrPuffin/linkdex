import { useEffect } from "react";

function Links(){
    function loadLinks(){
        console.info("Grabbed links...")
        fetch("http://127.0.0.1:8080/links")
            .then((response) => response.json())
            .then((data) => console.log(data));
    }

    useEffect(() => {loadLinks()});

    return <h1>link</h1>;
}

export default Links;