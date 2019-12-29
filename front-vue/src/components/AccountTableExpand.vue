<template>
  <div>
    <Row class="expand-row">
      <i-col span="8">
        <span class="expand-key">账户邮箱: </span>
        <span class="expand-value">{{ row.account_mail }}</span>
      </i-col>
      <i-col span="8">
        <span class="expand-key">账户密码: </span>
        <span class="expand-value">{{ row.account_password }}</span>
      </i-col>
      <i-col span="8">
        <span class="expand-key">验证邮箱: </span>
        <span class="expand-value">{{ row.verify_mail }}</span>
      </i-col>
    </Row>
    <Row>
      <i-col span="8">
        <span class="expand-key">创建日期: </span>
        <span class="expand-value">{{ row.create_date }}</span>
      </i-col>
      <i-col span="8">
        <span class="expand-key">授权日期: </span>
        <span class="expand-value">{{ row.agent_date }}</span>
      </i-col>
      <i-col span="8">
        <span class="expand-key">到期日期: </span>
        <span class="expand-value">{{ row.end_date }}</span>
      </i-col>
    </Row>
  </div>
</template>
<script>
export default {
  name: "ExpandRow",
  props: {
    row: Object
  }
};
</script>
<style scoped>
.expand-row {
  margin-bottom: 5px;
}
</style>
