import React, { useState } from 'react'
import VendorItem from "./VendorItem"

export default function Vendors(props) {

  return Object.entries(props.vendors).map(([key, value],index) => {
    return (
      <VendorItem
        key={index}
        title={key}
        image={value}
        valueVendorFilter={props.valueVendorFilter}
        setValueVendorFilter={props.setValueVendorFilter}
      />
    );
  });
}