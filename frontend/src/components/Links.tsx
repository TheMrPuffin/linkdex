import { Divider, Link } from "@mui/material";
import { Stack } from "@mui/system";
import { useEffect, useState } from "react";

function Links(){
    function loadLinks(){
        console.info("Grabbing links:")
        fetch("http://127.0.0.1:8080/links")
            .then((response) => response.json())
            .then((data) => {
                let linkObject = data.data.links
                const linksArray = linkObject.map((link: any) => 
                    <Link key={link.Id} href={link.url}>{link.name}</Link>);
                console.debug(linksArray);
                
                setLinks(linksArray)
            });
        }

    const [links, setLinks] = useState([]);
    useEffect(() => {loadLinks(); }, []);

    return <Stack sx={{ textAlign: 'center', color: 'primary.main'}} direction="column" divider={<Divider orientation="horizontal" flexItem />}
    spacing={2}>{links}</Stack>;
}

export default Links;