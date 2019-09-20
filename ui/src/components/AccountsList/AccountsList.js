import React from 'react'
import Table from 'react-bootstrap/Table'

const accounts = [
  {id: '0' ,name: 'a', available: 12.22},
  {id: '1' ,name: 'b', available: 992.22},
  {id: '2' ,name: 'c', available: 253.22}
]

export function AccountsList () {
  return <Table>
    <thead>
      <tr>
        <th>Nazwa</th>
        <th>Dostępne środki</th>
      </tr>
    </thead>
    <tbody>
      {accounts.map((account)=>
        <tr key={account.id}>
          <td>{account.name}</td>
          <td>{account.available}</td>
        </tr>
      )}
    </tbody>
  </Table>
}
