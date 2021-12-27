<template>
  <div>
    <date-picker 
    v-model="date"
    @change="setDate"
    @clear="delDate"
    ></date-picker>
    <vuetable ref="vuetable"
      :key="time"
      :api-url="getUrl()"
      :fields="fields"
      data-path="data"
      pagination-path=""
      :http-options="httpOptions"
      :per-page="15"
      @vuetable:pagination-data="onPaginationData"
    ></vuetable>

    <div class="pagination ui basic segment grid">
      <vuetable-pagination-info ref="paginationInfo"
      ></vuetable-pagination-info>
      
      <vuetable-pagination ref="pagination"
        @vuetable-pagination:change-page="onChangePage"
      ></vuetable-pagination>
    </div>
  </div>
</template>

<script>
import Vuetable from "vuetable-2/src/components/Vuetable";
import VuetablePagination from "vuetable-2/src/components/VuetablePagination";
import VuetablePaginationInfo from 'vuetable-2/src/components/VuetablePaginationInfo';

import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';

import Fields from './Fields.js'
import {readValue} from '../../actions/jwt.js'

let t = readValue()

export default {
  name: "UserList",
  components: {
    Vuetable,
    VuetablePagination,
    VuetablePaginationInfo,
    DatePicker
  },
  data() {
    return {
      fields: Fields,
      httpOptions: {
        headers: {
            "Authorization": "Bearer "+ t,
        },
        withCredentials: true 
      },
      time: 0,
      date: null,
    };
  },
  methods: {
    onPaginationData(paginationData) {
      this.$refs.pagination.setPaginationData(paginationData);
      this.$refs.paginationInfo.setPaginationData(paginationData);
    },
    onChangePage(page) {
      this.$refs.vuetable.changePage(page);
    },
    setDate(date){
      this.time = Date.parse(date) / 1000;
      this.$refs.vuetable.refresh();
    },
    delDate(){
      this.time = 0;
    },
    getUrl(){
      return "/api/user/list?" +
        "timestamp=" + this.time
    },
  }
};
</script>

<style>
  .pagination {
    margin-top: 1rem;
  }
  .ui.blue.table {
    border-top: 0.2em solid #5f9ea0;
  }
  .mx-input-wrapper{
    margin-bottom: 12px;
  }
  .mx-input:hover{
    border-color: #5f9ea0;
  }
  .mx-btn:hover {
    border-color: #5f9ea0;
    color: #75c1c4;
  }
  .mx-calendar-content .cell:hover {
    background-color: #9BC3C4;
    color:wheat
  }
  .mx-table-date .today {
    color: #75c1c4;
  }
  .mx-calendar-content .cell.active {
    color: wheat;
    background-color: #5f9ea0;
  }  
</style>