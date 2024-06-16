import {
  CitySelect,
  CountrySelect,
  StateSelect,
} from "react-country-state-city";
import "react-country-state-city/dist/react-country-state-city.css";

import React, { useState } from "react";


export default function LocationDropdown({ setLocation, location }) {
  const [countryid, setCountryid] = useState(0);
  const [stateid, setstateid] = useState(0);
  return (
    <div>
      <h6>Country</h6>
      <CountrySelect
        onChange={(e) => {
          setCountryid(e.id);

          setLocation(prevLocation => ({
            ...prevLocation,
            country: e.name,
          }));

          // debug 
          console.log("Country: ", e.name)
          
        }}
        placeHolder="Select Country"
      />

      <h6>State</h6>
      <StateSelect
        countryid={countryid}
        onChange={(e) => {
          setstateid(e.id);

          setLocation(prevLocation => ({
            ...prevLocation,
            state: e.name,
          }));

          // dubug
          console.log("State: ", e.name)
        }}
        placeHolder="Select State"
      />

      <h6>City</h6>
      <CitySelect
        countryid={countryid}
        stateid={stateid}
        onChange={(e) => {
          setLocation(prevLocation => ({
            ...prevLocation,
            city: e.name,
          }));
          
          // debug
          console.log("City: ", e.name)
        }}
        placeHolder="Select City"
      />
    </div>
  );
}