<template>
  <div>
    <vuetable ref="vuetable"
      api-url="http://localhost:8000/user/list"
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
import {readValue} from '../../actions/jwt.js'

let t = readValue()

export default {
  name: "UserList",
  components: {
    Vuetable,
    VuetablePagination,
    VuetablePaginationInfo
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
    };
  },
  methods: {
    onPaginationData(paginationData) {
      this.$refs.pagination.setPaginationData(paginationData);
      this.$refs.paginationInfo.setPaginationData(paginationData);
    },
    onChangePage(page) {
      this.$refs.vuetable.changePage(page);
    }
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
</style>