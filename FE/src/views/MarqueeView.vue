<template>
    <div class="q-pa-md">
      <q-table
        title="List Pengumuman"
        :rows="rows"
        :columns="columns"
        row-key="Id"
        dense
        class="my-sticky-dynamic"
        style="height: 80vh"
        virtual-scroll
        :rows-per-page-options="[0]"
        :filter="filter"
        :separator="separator"
        selection="single"
        v-model:selected="selected"
        :selected-rows-label="getSelectedRow"
      >
        <template #header-cell="props">
          <q-th :props="props" class="bg-primary glossy text-white text-h7">
            {{ props.col.label }}
          </q-th>
        </template>
  
        <template #body-cell-MarqueeText="props">
          <q-td :props="props" dense>
            <a href="#" @click="editRow(props)">{{ props.value }}</a>
          </q-td>
        </template>
  
        <template #top-right="props">
          <q-btn
            color="primary"
            style="margin: 3px"
            icon="add"
            dense
            @click="showForm"
            class="bg-primary glossy text-white"
          />
          <q-btn
            color="primary"
            style="margin: 3px"
            icon="archive"
            no-caps
            dense
            @click="exportTable"
          />
          <q-input
            outlined
            dense
            debounce="300"
            v-model="filter"
            placeholder="Search"
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
      </q-table>
  
  
      <q-dialog v-model="show_form">
        <q-card style="width: 600px; max-width: 60vw">
          <q-card-section>
            <q-btn
              round
              flat
              dense
              icon="close"
              class="float-right"
              color="grey-8"
              v-close-popup
            ></q-btn>
            <div class="text-h6">Add Row</div>
          </q-card-section>
          <q-separator inset></q-separator>
          <q-card-section class="q-pt-none">
            <q-form class="q-gutter-md">
              <q-list>
                <q-item>
                  <q-item-section>
                    <q-item-label class="q-pb-xs">Pengumuman</q-item-label>
                    <q-input dense outlined v-model="marquee.MarqueeText" />
                  </q-item-section>
                </q-item>                
              </q-list>
            </q-form>
          </q-card-section>
          <q-card-section>
            <q-card-actions align="right">
              <q-btn
                flat
                label="Cancel"
                color="warning"
                dense
                v-close-popup
              ></q-btn>
              <q-btn
                flat
                label="OK"
                color="primary"
                dense
                v-close-popup
                @click="updateRow"
              ></q-btn>
            </q-card-actions>
          </q-card-section>
        </q-card>
      </q-dialog>
  
      <q-dialog v-model="show_dialog">
        <q-card style="width: 600px; max-width: 30vw">
          <q-separator inset></q-separator>
          <q-card-section class="q-pt-none">
            <q-form class="q-gutter-md">
              <q-list>
                <q-item>
                  <q-item-section> Do you want to delete Record? </q-item-section>
                </q-item>
              </q-list>
            </q-form>
          </q-card-section>
          <q-card-section>
            <q-card-actions align="right">
              <q-btn
                flat
                label="Cancel"
                color="warning"
                dense
                v-close-popup
              ></q-btn>
              <q-btn
                flat
                label="OK"
                color="primary"
                dense
                v-close-popup
                @click="deleteRow"
              ></q-btn>
            </q-card-actions>
          </q-card-section>
        </q-card>
      </q-dialog>
    </div>
  </template>
  <script>
  import Get from "@/api/http-get";
  import Post from '@/api/http-post'
  import Update from '@/api/http-update'
  import Delete from "@/api/http-delete";
  import { Notify, exportFile } from "quasar";
  
  export default {
    data() {
      return {
        columns: [
          {
            name: "MarqueeText",
            label: "Isi Pengumuman",
            field: "MarqueeText",
            align: "left",
            sortable: true,
          },
        ],
        marquee : {
          Id: '',
          MarqueeText: '',
        },
        rows: [],
        filter: "",
        separator: "cell",
        selected: [],
        show_dialog: false,
        show_form: false
      };
    },
  
    async mounted() {
      this.getRow();
    },
  
    methods: {
      getSelectedRow() {
        this.show_dialog = true;
        this.marquee.Id = this.selected[0].Id
      },
  
      async getRow() {
        this.rows = (await Get.masjidMarquee()).data;
        // console.log('rows', this.rows)
      },
  
      editRow(props) {
        if (props) {
          this.marquee.Id = props.row.Id;
          this.marquee.MarqueeText = props.row.MarqueeText;
        }
  
        this.show_form = true;
      },
  
      async deleteRow() {
        // console.log('delete', props.row.Id)
        const id = this.marquee.Id;
  
        // const index = this.rows.indexOf(id);
        await Delete.masjidMarquee(id).then((res) => {
          let status = res.data.status;
          let message = res.data.message;
          if (status) {
            Notify.create({
              message,
              color: "blue-8",
              textColor: "white",
              icon: "cloud_done",
            });      
          }
          
        });
        
        await this.getRow();
        this.show_dialog = false
      },
  
      async updateRow() {
        // console.log("update row");
        let id = this.marquee.Id;
        // console.log("id", id);
        // console.log('role', this.role)
  
        if (id > 0) {
          await Update.masjidMarquee(id, this.marquee).then((res) => {
            // console.log('res', res)
            // console.log('status', res.data.status)
            // console.log('message', res.data.message)
  
            const status = res.data.status;
            const message = res.data.message;
            // console.log("after update", status);
  
            if (status) {
              Notify.create({
                message,
                color: "blue-8",
                textColor: "white",
                icon: "cloud_done",
              });
            }
          });
        } else {
          let data = {
            MarqueeText: this.marquee.MarqueeText,
          };
  
        //   console.log('data', data)
  
          await Post.masjidMarquee(data).then((res) => {
            const status = res.data.status;
            const message = res.data.message;
            // console.log("after update", status);
  
            if (status) {
              Notify.create({
                message,
                color: "blue-8",
                textColor: "white",
                icon: "cloud_done",
              });
            }
          });
        }
  
        await this.getRow();
        this.show_dialog = false
      },
  
  
      showForm() {
        // console.log("add");
        // this.$router.push({ name: "slidesForm" });
        this.marquee.Id = "";
        this.marquee.MarqueeText = "";
        this.show_form = true;
      },
  
      wrapCsvValue(val, formatFn, row) {
        let formatted = formatFn !== void 0 ? formatFn(val, row) : val;
  
        formatted =
          formatted === void 0 || formatted === null ? "" : String(formatted);
  
        formatted = formatted.split('"').join('""');
        /**
         * Excel accepts \n and \r in strings, but some other CSV parsers do not
         * Uncomment the next two lines to escape new lines
         */
        // .split('\n').join('\\n')
        // .split('\r').join('\\r')
  
        return `"${formatted}"`;
      },
  
      exportTable() {
        // naive encoding to csv format
        const content = [this.columns.map((col) => this.wrapCsvValue(col.label))]
          .concat(
            this.rows.map((row) =>
              this.columns
                .map((col) =>
                  this.wrapCsvValue(
                    typeof col.field === "function"
                      ? col.field(row)
                      : row[col.field === void 0 ? col.name : col.field],
                    col.format,
                    row
                  )
                )
                .join(",")
            )
          )
          .join("\r\n");
  
        const status = exportFile("table-export.csv", content, "text/csv");
  
        if (status !== true) {
          $q.notify({
            message: "Browser denied file download...",
            color: "negative",
            icon: "warning",
          });
        }
      },
    },
  };
  </script>
  <style>
  .my-sticky-dynamic {
    /* height or max-height is important */
    height: 410px;
    /* this will be the loading indicator */
    /* prevent scrolling behind sticky top row on focus */
  }
  .my-sticky-dynamic .q-table__top,
  .my-sticky-dynamic .q-table__bottom,
  .my-sticky-dynamic thead tr:first-child th {
    background-color: white;
  }
  .my-sticky-dynamic thead tr th {
    position: sticky;
    z-index: 1;
  }
  .my-sticky-dynamic thead tr:last-child th {
    /* height of all previous header rows */
    top: 48px;
  }
  .my-sticky-dynamic thead tr:first-child th {
    top: 0;
  }
  .my-sticky-dynamic tbody {
    /* height of all previous header rows */
    scroll-margin-top: 48px;
  }
  </style>
  