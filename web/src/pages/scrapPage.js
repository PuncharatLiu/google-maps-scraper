import React, { useState } from "react";
import axios from "axios";

function ScrapPage() {
  const [data, setData] = useState(null);
  const [keyword, setKeyword] = useState(null);

  return (
    <div>
      <KeyInput keyword={keyword} setKeyword={setKeyword} />
      <ScrapeButton setData={setData} keyword={keyword} />
      {data && <CrefAndUpload />}
      <CSVTable data={data} />
    </div>
  );
}

function KeyInput({ keyword, setKeyword }) {
  // const [key, setKey] = useState("");
  const handleKeyChange = (event) => {
    setKeyword(event.target.value);
    console.log("keyword: ", keyword);
  };

  return (
    <div>
      <input placeholder="Type location or key" value={keyword} onChange={handleKeyChange} />
    </div>
  );
}

function ScrapeButton({ setData, keyword }) {
  const handleScrap = async () => {
    try {
      const sendData = {
        key: keyword,
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
      </table>
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
