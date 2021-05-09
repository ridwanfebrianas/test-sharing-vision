import React, { useState, useEffect } from 'react'
import { useHistory, useLocation } from 'react-router-dom'
import {
  CCard,
  CCardBody,
  CCardHeader,
  CCol,
  CDataTable,
  CRow,
  CPagination
} from '@coreui/react'

import axios from 'axios';
var jsonData = []

const UserManagement = () => {
  const history = useHistory()
  const queryPage = useLocation().search.match(/page=([0-9]+)/, '')
  const currentPage = Number(queryPage && queryPage[1] ? queryPage[1] : 1)
  const [page, setPage] = useState(currentPage)
  const url = 'http://localhost:8080/user';

  
  async function axiosTest() {
    const response = await axios.get(url)
    //console.log('data yang didapat',response.data)
    return response.data
  }

  jsonData = axiosTest()
  //var dummy = []

  const pageChange = newPage => {
    currentPage !== newPage && history.push(`/users?page=${newPage}`)
  }

  useEffect(() => {
    currentPage !== page && setPage(currentPage)
  }, [currentPage, page])

  console.log('data', jsonData)
  return (
    <CRow>
      <CCol xl={6}>
        <CCard>
          <CCardHeader>
            Users List
          </CCardHeader>
          <CCardBody>
          <CDataTable
            items={jsonData}
            fields={[
              { key: 'id', _classes: 'font-weight-bold' },
              'name', 'username', 'password'
            ]}
            hover
            striped
            itemsPerPage={5}
            activePage={page}
            clickableRows
            onRowClick={(item) => history.push(`/users/${item.id}`)}
          />
          <CPagination
            activePage={page}
            onActivePageChange={pageChange}
            pages={5}
            doubleArrows={false} 
            align="center"
          />
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  )
}

export default UserManagement
