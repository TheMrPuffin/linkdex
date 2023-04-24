import { useEffect, useState } from "react";

function Links(){
    function loadLinks(){
        console.info("Grabbing links:")
        fetch("http://127.0.0.1:8080/links")
            .then((response) => response.json())
            .then((data) => {
                let linkObject = data.data.links
                const linksArray = linkObject.map((link: any) => 
                    <li key={link.Id}> {link.name}</li>);
                
                console.debug(linksArray);
                
                setLinks(linksArray)
            });
        }

    const [links, setLinks] = useState([]);
    useEffect(() => {loadLinks(); }, []);

    return <ul>{links}</ul>;
}

export default Links;