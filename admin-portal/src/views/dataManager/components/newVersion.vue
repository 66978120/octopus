<template>
  <div>
    <el-dialog title="创建新版本" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false" :show-close="close">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="true" />
        </el-form-item>
        <el-form-item label="版本描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :autosize="{ minRows: 2, maxRows: 4}" placeholder="请输入数据集描述" maxlength="300"
            show-word-limit />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button v-show="!showUpload" type="text" @click="nextStep('ruleForm')">下一步</el-button>
        </el-form-item>
        <el-form-item v-if="showUpload" label="数据集上传" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel"
            @upload="isCloseX" />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  import upload from '@/components/upload/index.vue'
  import { createNewVersion } from "@/api/dataManager.js";
  export default {
    name: "NewVersion",
    components: {
      upload
    },
    props: {
      row: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        showUpload: false,
        uploadData: { data: {}, type: undefined },
        ruleForm: {
          desc: ""
        },
        rules: {
          path: [
            {
              required: true,
              message: "请上传数据集",
              trigger: "change"
            }
          ]
        },
        CreateFormVisible: true,
        formLabelWidth: "120px",
        close: true
      }
    },
    created() {
      const { name, type } = this.row
      this.ruleForm = { name, type }
    },
    methods: {
      handleDialogClose() {
        this.$emit("close", false);
      },
      nextStep(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.showUpload = true
            const param = {
              desc: this.ruleForm.desc,
              datasetId: this.row.id
            }
            createNewVersion(param).then(response => {
              if (response.success) {
                this.uploadData.type = "newPreDatasetVersion"
                this.uploadData.datasetId = response.data.datasetId
                this.uploadData.version = response.data.version
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
              }
            })
          } else {
            return false;
          }
        });
      },
      cancel() {
        this.$emit("cancel", false);
      },
      confirm(val) {
        this.$emit("confirm", val);
      },
      isCloseX(val) {
        this.close = val
      }
    }
  }
</script>