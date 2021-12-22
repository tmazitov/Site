<template>
  <div>
    <date-picker 
    v-model="date" 
    type="date"
    @change="setDate"
    @clear="delDate"
    >
    </date-picker>
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
import Fields from './Fields.js'
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
import {readValue} from '../../actions/jwt.js'

let t = readValue()

export default {
  name: "UserList",
  components: {
    Vuetable,
    DatePicker,
    VuetablePagination,
    VuetablePaginationInfo,
  },
  data() {
    return {
      fields: Fields,
      time: 0,
      httpOptions: {
        headers: {
            "Authorization": "Bearer "+ t,
        },
        withCredentials: true 
      },
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
      return "http://localhost:8000/user/list?" +
        "t=" + this.time
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
  .mx-datepicker{
    margin-bottom: 12px;
  }
  .mx-input:hover{
    border-color: #588C8E;
  }

  .mx-table-date .today {
    color: #588C8E;
  }

  .mx-calendar-content .cell:hover {
    color: #73879c;
    background-color: #bad5d6;
  }

  .mx-calendar-content .cell.active{
    background-color: #588C8E;
    color: wheat;
  }
  .mx-btn:hover {
    border-color: #588C8E;
    color: #588C8E;
}
</style>