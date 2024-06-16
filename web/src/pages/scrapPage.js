import React, { useState, useEffect } from "react";
import axios from "axios";

import LocationDropdown from "../components/scrape_page/location_dropdown.js";

function ScrapPage() {
  const [data, setData] = useState(null);
 
  const [keyword, setKeyword] = useState(null);  
  const [location, setLocation] = useState({
    country: "",
    state: "",
    city: ""
  })

  const [address, setAddress] = useState("")
  const [isLoading, setIsLoading] = useState(false)

  
  useEffect(() => {
    const concatAddress = `${keyword} in ${location.city}, ${location.state}, ${location.country}`;
    setAddress(concatAddress);
    
    // debug
    console.log("address: ", concatAddress);
  }, [keyword, location]);
  
  

  return (
    <div>
      <KeyInput keyword={keyword} setKeyword={setKeyword} />
      <LocationDropdown setLocation={setLocation} location={location}/>
      <ScrapeButton setData={setData} address={address} setIsLoading={setIsLoading} />
      
      {data && <CrefAndUpload />}
      
      {isLoading && !data && <div>Scraping...</div>}
      {!isLoading && !data && <div>No data available. Please click the button to scrape data.</div>}
      {data && <CSVTable data={data} />}
    </div>
  );
}

function KeyInput({ keyword, setKeyword }) {
  const handleKeyChange = (event) => {
    setKeyword(event.target.value);
  };

  return (
    <div>
      <input placeholder="Type location or key" value={keyword} onChange={handleKeyChange} />
    </div>
  );
}

function ScrapeButton({ setData, address, setIsLoading }) {
  const handleScrap = async () => {
    setIsLoading(true)

    try {
      const sendData = {
        key: address,
        name: "result.csv",
      };

      const response = await axios.post('http://localhost:8080/scrape', sendData, {
        responseType: 'json',
      });

      // Debug
      console.log("response: ", response.data);
      setData(response.data); // Set the received data
    } catch (error) {
      console.error("Error fetching data: ", error);
    }
  };

  return (
    <button onClick={handleScrap}>Scrap</button>        
  );
}

function CSVTable({ data }) {
  if (!data) {
    return <div>No data available</div>;
  }

  const { titles, address, website, phone, Cid } = data;

  if (!titles || !address || !website || !phone || !Cid) {
    return <div>Data is incomplete</div>;
  }

  const tableData = [];

  for (let i = 1; i < titles.length; i++) {
    tableData.push({
      cid: Cid[i],
      title: titles[i],
      address: address[i],
      website: website[i],
      phone: phone[i],
    });
  }

  return (
    <div>      
      <table border="1">
        <thead>
          <tr>
            <th>CID</th>
            <th>Title</th>
            <th>Address</th>
            <th>Website</th>
            <th>Phone</th>
          </tr>
        </thead>
        <tbody>
          {tableData.map((row, index) => (
            <tr key={index}>
              <td>{row.cid}</td>
              <td>{row.title}</td>
              <td>{row.address}</td>
              <td>
                {row.website ? (
                  <a href={row.website} target="_blank" rel="noopener noreferrer">
                    {row.website}
                  </a>
                ) : (
                  "N/A"
                )}
              </td>
              <td>{row.phone}</td>
            </tr>
          ))}
        </tbody>
      </table>l
    </div>
  );
}

function CrefAndUpload() {
  const handleCrefAndUpload = async () => {
    await axios.post('http://localhost:8080/cref');
  }

  return (
    <div>
      <button onClick={handleCrefAndUpload}>Cross Reference and Upload</button>
    </div>
  )
}

export default ScrapPage;
