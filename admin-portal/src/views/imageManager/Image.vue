<template>
  <div>
    <searchForm
      :search-form="searchForm"
      class="searchForm"
      :blur-name="'镜像名称/标签/描述 搜索'"
      @searchData="getSearchData"
    />
    <el-button v-if="!flag" type="primary" class="create" @click="create"
      >创建</el-button
    >
    <el-table
      :data="tableData"
      style="width: 100%; font-size: 15px"
      :header-cell-style="{ 'text-align': 'left', color: 'black' }"
      :cell-style="{ 'text-align': 'left' }"
    >
      <el-table-column label="镜像名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.imageName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像标签" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.imageVersion }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="镜像描述"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.imageDesc }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="flag" label="群组名">
        <template slot-scope="scope">
          <span>{{
            scope.row.spaceName === "" ? "默认群组" : scope.row.spaceName
          }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="flag" label="提供者" align="center">
        <template slot-scope="scope">
          <el-tooltip
            trigger="hover"
            :content="scope.row.userEmail"
            placement="top"
          >
            <span>{{ scope.row.username }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        label="镜像地址"
        align="center"
        :show-overflow-tooltip="true"
      >
        <template slot-scope="scope">
          <span>{{
            scope.row.sourceType === 2 ? scope.row.imageAddr : ""
          }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像来源" align="center">
        <template slot-scope="scope">
          <span>{{ sourceType(scope.row.sourceType) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.createdAt | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center">
        <template slot-scope="scope">
          <span v-if="!(scope.row.progress && scope.row.progress != 0)">{{
            imageStatus(scope.row.imageStatus)
          }}</span>
          <span v-if="scope.row.progress && scope.row.progress != 0">{{
            "上传中"
          }}</span>
          <el-progress
            :percentage="parseInt(scope.row.progress - 1)"
            v-if="scope.row.progress && scope.row.progress != 0"
          ></el-progress>
        </template>
      </el-table-column>
      <el-table-column v-if="!flag" label="操作" align="center" width="250">
        <template slot-scope="scope">
          <el-button
            v-if="scope.row.imageStatus == 1 || scope.row.imageStatus == 4"
            type="text"
            @click="handleEdit(scope.row)"
            :disabled="scope.row.progress && scope.row.progress != 0"
            >重新上传
          </el-button>
          <el-button
            type="text"
            @click="open2(scope.row)"
            :disabled="scope.row.progress && scope.row.progress != 0"
            >删除</el-button
          >
          <el-button type="text" @click="open(scope.row)">修改描述</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination
        :current-page="searchData.pageIndex"
        :page-sizes="[10, 20, 50, 80]"
        :page-size="searchData.pageSize"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <!-- 镜像对话框 -->
    <dialogForm
      v-if="FormVisible"
      :flag="Logo"
      :row="row"
      @cancel="cancel"
      @confirm="confirm"
      @close="close"
    />
  </div>
</template>
<script>
import dialogForm from "./components/dialogForm.vue";
import {
  getUserImage,
  getPreImage,
  deletePreImage,
  editPreImage,
} from "@/api/imageManager.js";
// import { groupDetail } from '@/api/userManager.js'
import searchForm from "@/components/search/index.vue";
import store from "@/store";
export default {
  name: "PreImage",
  components: {
    dialogForm,
    searchForm,
  },
  props: {
    imageTabType: { type: Number, default: undefined },
  },
  data() {
    return {
      tableData: [],
      row: {},
      total: undefined,
      FormVisible: false,
      flag: true,
      Logo: true,
      searchForm: [
        {
          type: "Select",
          label: "状态",
          prop: "imageStatus",
          placeholder: "请选择状态",
          options: [
            { label: "未制作", value: 1 },
            { label: "制作中", value: 2 },
            { label: "制作完成", value: 3 },
            { label: "制作失败", value: 4 },
          ],
        },
        {
          type: "Input",
          label: "镜像名",
          prop: "imageNameLike",
          placeholder: "请输入镜像名",
        },
        {
          type: "Select",
          label: "来源类型",
          prop: "sourceType",
          placeholder: "请选择来源类型",
          options: [
            { label: "上传", value: 1 },
            { label: "远程", value: 2 },
            { label: "保存", value: 3 },
          ],
        },
      ],
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      },
    };
  },
  created() {
    this.getImage(this.searchData);
    // this.timer = setInterval(() => { this.getImage(this.searchData) }, 1000)
    if (this.imageTabType !== 1) {
      this.flag = false;
      this.timer = setInterval(() => {
        this.getImage(this.searchData);
      }, 2000);
      this.searchForm[2].options = [
        { label: "上传", value: 1 },
        { label: "远程", value: 2 },
      ];
    } else {
      this.searchForm.push(
        {
          type: "InputSelectUser",
          label: "用户",
          prop: "userId",
          placeholder: "请输入用户名",
        },
        {
          type: "InputSelectGroup",
          label: "群组",
          prop: "spaceId",
          placeholder: "请输入群组名",
        }
      );
      // this.timer = setInterval(() => { this.getImage(this.searchData) }, 1000)
    }
  },
  mounted() {
    window.addEventListener("beforeunload", (e) => {
      sessionStorage.clear();
    });
  },
  destroyed() {
    window.removeEventListener("beforeunload", (e) => {
      sessionStorage.clear();
    });
    clearInterval(this.timer);
    this.timer = null;
  },
  methods: {
    getImage(data) {
      this.type = this.imageTabType;
      if (this.type === 1) {
        getUserImage(data).then((response) => {
          if (response.success) {
            this.tableData = response.data.images;
            this.total = response.data.totalSize;
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: "warning",
            });
          }
        });
      }
      if (this.type === 2) {
        getPreImage(data).then((response) => {
          if (response.success) {
            if (response.data !== null) {
              this.total = parseInt(response.data.totalSize);
              (this.tableData = response.data.images),
                this.tableData.forEach((item) => {
                  if (sessionStorage.getItem(JSON.stringify(item.id))) {
                    item.progress = sessionStorage.getItem(
                      JSON.stringify(item.id)
                    );
                  }
                });
            }
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: "warning",
            });
          }
        });
      }
    },
    handleEdit(row) {
      this.row = row;
      this.FormVisible = true;
      this.Logo = false;
      store.commit("user/SET_PROGRESSID", row.id);
    },
    handleDelete(row) {
      deletePreImage(row.id).then((response) => {
        if (response.success) {
          this.$message({
            message: "删除成功",
            type: "success",
          });
          this.getImage(this.searchData);
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: "warning",
          });
        }
      });
    },
    handleSizeChange(val) {
      this.searchData.pageSize = val;
      this.getImage(this.searchData);
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val;
      this.getImage(this.searchData);
    },
    cancel(val) {
      this.FormVisible = val;
      this.getImage(this.searchData);
    },
    confirm(val) {
      this.FormVisible = val;
      this.getImage(this.searchData);
    },
    close(val) {
      this.FormVisible = val;
      this.getImage(this.searchData);
    },
    create() {
      this.FormVisible = true;
      this.row = {};
      this.Logo = true;
    },
    getSearchData(val) {
      this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize };
      this.searchData = Object.assign(val, this.searchData);
      if (
        this.searchData.spaceNameLike &&
        this.searchData.spaceNameLike == "默认群组"
      ) {
        this.searchData.spaceNameLike = "";
      }
      this.getImage(this.searchData);
    },
    // 镜像状态
    imageStatus(value) {
      switch (value) {
        case 1:
          return "未上传";
        case 2:
          return "制作中";
        case 3:
          return "制作完成";
        case 4:
          return "制作失败";
      }
    },
    sourceType(value) {
      switch (value) {
        case 1:
          return "上传";
        case 2:
          return "远程";
        default:
          return "保存";
      }
    },
    // 修改描述
    open(val) {
      const data = val;
      this.$prompt("请输入描述", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          editPreImage({
            id: data.id,
            imageName: data.imageName,
            imageVersion: data.imageVersion,
            imageType: data.imageType,
            imageAddr: data.imageAddr,
            imageDesc: value,
          }).then((response) => {
            if (response.success) {
              this.$message({
                message: "编辑描述成功",
                type: "success",
              });
              this.getImage(this.searchData);
              this.$emit("confirm", false);
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: "warning",
              });
            }
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
    // 删除确认
    open2(val) {
      this.$confirm("此操作将永久删除该镜像, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.handleDelete(val);
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
  },
};
</script>
<style lang="scss" scoped>
.create {
  float: right;
}

.block {
  float: right;
  margin: 20px;
}

.searchForm {
  display: inline-block;
}
</style>