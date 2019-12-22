<template lang="pug">
  el-container#app
    el-header
      el-dropdown
        span.el-dropdown-link 選單
          i.el-icon-setting.el-icon--right
        el-dropdown-menu(slot='dropdown')
          el-dropdown-item
            el-button(type="text" @click="handleCreate") 新增
    el-main
      el-table(v-loading="loading" :data='endpoint' style='width: 100%' stripe)
        el-table-column(prop='name' label='名稱' width='180' sortable)
        el-table-column(label='網址' sortable)
          template(slot-scope="scope")
            el-link(type="primary" :href="scope.row.url" target="_blank") {{ scope.row.url }}
        el-table-column(label='狀態碼' width='180' sortable)
          template(slot-scope="scope")
            el-tag(size="medium" :type="getTagType(scope.row.status)") {{ !scope.row.status || scope.row.status === -1 ? 'unknow' : scope.row.status }}
        el-table-column(label='更新時間')
          template(slot-scope="scope")
            i.el-icon-time
            span(style="margin-left: 10px") {{ new Date(scope.row.updated_at*1000) }}
        el-table-column(width='180')
          template(slot-scope="scope")
            el-button(size="mini" @click="handleEdit(scope.row)") 編輯
            el-button(size="mini" type="danger" @click="handleDelete(scope.row)") 刪除
    el-dialog(title="新增" :visible.sync="createDialogForm.visible" width="50%")
      el-form(:model="createDialogForm" label-width="80px")
        el-form-item(label="名稱")
          el-input(v-model="createDialogForm.name")
        el-form-item(label="網址")
          el-input(v-model="createDialogForm.url")
      span.dialog-footer(slot="footer")
        el-button(@click="createDialogForm.visible = false") 取消
        el-button(type="primary" @click="handleCreateConfirm") 確定
    el-dialog(title="編輯" :visible.sync="editDialogForm.visible" width="50%")
      el-form(:model="editDialogForm" label-width="80px")
        el-form-item(label="名稱")
          el-input(v-model="editDialogForm.name")
        el-form-item(label="網址")
          el-input(v-model="editDialogForm.url")
      span.dialog-footer(slot="footer")
        el-button(@click="editDialogForm.visible = false") 取消
        el-button(type="primary" @click="handleEditConfirm") 確定
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      loading: true,
      createDialogForm: {
        visible: false,
        name: "",
        url: ""
      },
      editDialogForm: {
        visible: false,
        id: 0,
        name: "",
        url: ""
      },
      endpoint: []
    };
  },
  async mounted() {
    this.loading = true;
    await this.loadEndpoint();
    this.loading = false;
    setInterval(this.loadEndpoint, 5000);
  },
  methods: {
    async loadEndpoint() {
      const res = await this.$axios.get("/api/endpoint");
      this.endpoint = res.data;
    },
    getTagType(status) {
      if (!status || status === -1) {
        return "warning";
      } else if (status <= 299) {
        return "success";
      } else {
        return "danger";
      }
    },
    handleCreate() {
      this.createDialogForm.name = "";
      this.createDialogForm.url = "";
      this.createDialogForm.visible = true;
    },
    async handleCreateConfirm() {
      await this.$axios.post("/api/endpoint", this.createDialogForm);
      this.createDialogForm.visible = false;
      await this.loadEndpoint();
    },
    handleEdit(row) {
      this.editDialogForm.id = row.id;
      this.editDialogForm.name = row.name;
      this.editDialogForm.url = row.url;
      this.editDialogForm.visible = true;
    },
    async handleEditConfirm() {
      await this.$axios.put(
        `/api/endpoint/${this.editDialogForm.id}`,
        this.editDialogForm
      );
      this.editDialogForm.visible = false;
      await this.loadEndpoint();
    },
    async handleDelete(row) {
      await this.$axios.delete(`/api/endpoint/${row.id}`);
      await this.loadEndpoint();
    }
  }
};
</script>

<style lang="sass">
 body
   margin: 0
 #app
   font-family: "Avenir", Helvetica, Arial, sans-serif
   -webkit-font-smoothing: antialiased
   -moz-osx-font-smoothing: grayscale
   color: #2c3e50

.el-header
   background-color: #B3C0D1
   color: #333
   line-height: 60px
   text-align: right
   font-size: 12px
</style>
