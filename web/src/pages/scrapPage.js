import React, { useState } from "react";
import axios from "axios";

function ScrapPage() {
  const [data, setData] = useState(null); // Initialize data with null

  return (
    <div>
      <KeyInput />
      <ScrapeButton setData={setData} />
      <CSVTable data={data} />
    </div>
  );
}

function KeyInput() {
  const [key, setKey] = useState("");

  const handleKeyChange = (event) => {
    setKey(event.target.value);
    console.log("key: ", key);
  };

  return (
    <div>
      <input placeholder="Type location or key" value={key} onChange={handleKeyChange} />
    </div>
  );
}

function ScrapeButton({ setData }) {
  const handleScrap = async () => {
    try {
      const sendData = {
        key: "Korean restaurant in Bangkok",
        name: "test_result.csv",
      };
      const response = await axios.post('http://localhost:8080/scrape', sendData, {
        responseType: 'json', // Ensure the response type is JSON
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
      <h1>Korean Restaurants in Bangkok</h1>
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
      </table>
    </div>
  );
}

export default ScrapPage;
