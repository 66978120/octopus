<template>
  <div>
    <el-dialog :close-on-click-modal="false" :title="title" width="70%" :visible.sync="versionListVisible"
      :before-close="handleDialogClose">
      <el-table v-loading="loading" :data="versionList" style="width: 100%" height="350">
        <el-table-column label="版本号">
          <template slot-scope="scope">
            <span>{{ scope.row.algorithmVersion }}</span>
          </template>
        </el-table-column>
        <el-table-column label="版本描述" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ scope.row.algorithmDescript }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ scope.row.createdAt | parseTime }}</span>
          </template>
        </el-table-column>
        <el-table-column v-if="algorithmTabType === 1 ? false :true" label="提供者">
          <template slot-scope="scope">
            <span>{{ scope.row.userName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="算法状态">
          <template slot-scope="scope">
            <span v-if="!(scope.row.progress&&scope.row.progress!=0)">{{ getAlgorithmStatus(scope.row.fileStatus)
              }}</span>
            <span v-if="scope.row.progress&&scope.row.progress!=0">{{ "上传中" }}</span>
            <el-progress :percentage="parseInt(scope.row.progress-1)" v-if="scope.row.progress&&scope.row.progress!=0">
            </el-progress>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <!-- <el-button type="text">预览</el-button> -->
            <el-button v-show="algorithmTabType === 1 ? true :false"
              v-if="(scope.row.fileStatus === 1 ) || (scope.row.fileStatus === 4 ) ? true : false" type="text"
              @click="reupload(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">重新上传
            </el-button>
            <el-button type="text" style="padding-right:10px" :disabled="(scope.row.fileStatus === 3)? false : true"
              @click="createTask(scope.row)">
              创建训练任务
            </el-button>
            <el-button slot="reference" type="text" :disabled="scope.row.fileStatus === 3 ? false : true"
              @click="confirmDownload(scope.row)">
              下载
            </el-button>
            <el-button v-if="algorithmTabType === 1 ? true :false" slot="reference" style="padding-right:10px"
              type="text" :disabled="scope.row.fileStatus === 3 ? false : true" @click="confirmShare(scope.row)">
              {{ scope.row.isShared?"取消分享":"分享" }}
            </el-button>
            <el-button v-if="algorithmTabType === 1 ? true :false" slot="reference" type="text"
              @click="confirmDelete(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
          @current-change="handleCurrentChange" />
      </div>
      <div slot="footer">
      </div>
    </el-dialog>
    <reuploadAlgorithm v-if="myAlgorithmVisible" :reupload-data="reuploadData" @close="close" @cancel="cancel"
      @confirm="confirm" />
  </div>
</template>

<script>
  import { getPubAlgorithmVersionList, getAlgorithmVersionList, queryAlgorithmVersion, compressAlgorithm, downloadAlgorithm, shareAlgorithmVersion, cancelShareAlgorithmVersion, deleteAlgorithmVersion } from "@/api/modelDev"
  import reuploadAlgorithm from './reuploadAlgorithm.vue'
  import store from '@/store'
  export default {
    name: "VersionList",
    components: {
      reuploadAlgorithm
    },
    props: {
      algorithmTabType: { type: Number, default: undefined },
      data: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        title: '版本列表/' + this.data.algorithmName,
        versionListVisible: true,
        myAlgorithmVisible: false,
        loading: false,
        pageIndex: 1,
        pageSize: 20,
        total: undefined,
        typeChange: undefined,
        versionList: [],
        shareTitle: "是否分享至本群组，分享后群内所有人员可见",
        timer: null,
        reuploadData: {}
      }
    },
    created() {
      this.getVersionList()
      this.timer = setInterval(() => { this.getVersionList() }, 2000)

    },
    destroyed() {
      clearInterval(this.timer)
      this.timer = null
    },
    methods: {
      reupload(row) {
        store.commit('user/SET_PROGRESSID', row.algorithmId + row.algorithmVersion)
        this.myAlgorithmVisible = true
        this.reuploadData = row
      },
      handleSizeChange(val) {
        this.pageSize = val
        this.getVersionList()
      },
      handleCurrentChange(val) {
        this.pageIndex = val
        this.getVersionList()
      },
      getVersionList(param) {
        this.typeChange = this.algorithmTabType
        if (!param) {
          param = { pageIndex: this.pageIndex, pageSize: this.pageSize }
        }
        param.algorithmId = this.data.algorithmId
        if (this.typeChange === 2) {
          getPubAlgorithmVersionList(param).then(response => {
            if (response.success) {
              this.versionList = response.data.algorithms
              this.total = response.data.totalSize
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          getAlgorithmVersionList(param).then(response => {
            if (response.success) {
              const newArr = []
              response.data.algorithms.filter(function (item, index) {
                const obj = item.algorithmDetail
                obj.isShared = item.isShared
                newArr.push(obj)
              })
              this.versionList = newArr
              this.versionList.forEach(item => {
                if (sessionStorage.getItem(JSON.stringify(item.algorithmId + item.algorithmVersion))) {
                  item.progress = sessionStorage.getItem(JSON.stringify(item.algorithmId + item.algorithmVersion))
                }

              })
              this.total = response.data.totalSize
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        }
      },
      // 接受到url下载
      URLdownload(fileName, url) {
        const link = document.createElement('a')
        link.style = 'display: none'; // 创建一个隐藏的a标签
        link.setAttribute('download', fileName)
        link.setAttribute('href', url)
        link.setAttribute('target', "_blank")
        document.body.appendChild(link);
        link.click(); // 触发a标签的click事件
        document.body.removeChild(link);
      },
      confirmDownload(row) {
        this.$confirm('是否下载此版本？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.handleDownload(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleDownload(row) {
        const that = this
        this.loading = true
        const param = {
          algorithmId: row.algorithmId,
          version: row.algorithmVersion
        }
        let latestCompressed = row.latestCompressed
        compressAlgorithm(param).then(response => {
          if (response.success) {
            param.compressAt = response.data.compressAt
            param.domain = this.GLOBAL.DOMAIN
            const interval = setInterval(function () {
              queryAlgorithmVersion(param).then(response => {
                if (response.success) {
                  latestCompressed = response.data.algorithm.latestCompressed
                } else {
                  that.loading = false
                  clearInterval(interval)
                  that.$message({
                    message: that.getErrorMsg(response.error.subcode),
                    type: 'warning'
                  });
                }
              })
              if (param.compressAt <= latestCompressed) {
                that.loading = false
                clearInterval(interval)
                downloadAlgorithm(param).then(response => {
                  if (response.success) {
                    that.URLdownload(row.algorithmName, response.data.downloadUrl)
                    that.$message.success("下载成功");
                  } else {
                    that.$message({
                      message: that.getErrorMsg(response.error.subcode),
                      type: 'warning'
                    });
                  }
                })
              }
            }, 3000)
          } else {
            that.loading = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      handleDialogClose() {
        this.$emit("close", false);
      },
      confirmShare(row) {
        if (row.isShared > 0) {
          this.shareTitle = "是否取消本群组分享？"
        } else {
          this.shareTitle = "是否分享至本群组，分享后群内所有人员可见"
        }
        this.$confirm(this.shareTitle, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.handleShare(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleShare(row) {
        const param = {
          algorithmId: row.algorithmId,
          version: row.algorithmVersion
        }
        if (row.isShared) {
          cancelShareAlgorithmVersion(param).then(response => {
            if (response.success) {
              this.$message.success("已取消本群组分享");
              this.getVersionList();
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          shareAlgorithmVersion(param).then(response => {
            if (response.success) {
              this.$message.success("已分享至群组");
              this.getVersionList();
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        }
      },
      getAlgorithmStatus(value) {
        switch (value) {
          case 1:
            return "未上传"
          case 2:
            return "上传中"
          case 3:
            return "解压完成"
          case 4:
            return "解压失败"
        }
      },
      cancel(val) {
        this.algorithmVersionDeleteDialogVisible = val
        this.myAlgorithmVisible = val
      },
      close(val) {
        this.algorithmVersionDeleteDialogVisible = val
        this.myAlgorithmVisible = val
      },
      confirm(val) {
        this.myAlgorithmVisible = val
      },
      confirmDelete(row) {
        this.$confirm('此操作将永久删除此算法版本(如该版本已分享，则分享版本也会被删除)，是否继续', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.handleDelete(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleDelete(row) {
        deleteAlgorithmVersion(row).then(response => {
          if (response.success) {
            this.$message.success("删除成功");
            this.getVersionList();
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      // 创建训练任务
      createTask(row) {
        const data = row
        data.trainingTask = true
        // data.open = true
        switch (this.algorithmTabType) {
          case 1:
            data.type = '我的算法'
            break;
          case 2:
            data.type = '公共算法'
            break;
          default:
            data.type = '预置算法'
        }
        this.$router.push({
          name: 'trainingManager',
          params: { data: data }
        })
      }
    }
  }
</script>
<style lang="scss" scoped>
  .pagination {
    float: right;
    margin: 20px;
  }
</style>