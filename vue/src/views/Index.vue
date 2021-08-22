<template>
	<div>
		<div class="container">
			<!-- <SpWelcome /> -->
			<div style="height: 200px; padding: 1px 0 0 0;">
  <div class="card" style="background-color:white; padding:10px; border-radius:8px;">
    <h2 class="cardtitle">Orin Pay</h2>
    <div class="cardcontent" style="color:#ggg">Bridging crypto to the masses</div>
  </div>
</div>
			<SpTokenSend :address="address" />
			<SpTransferList :address="address" />
		</div>
	</div>
</template>

<script>
export default {
	name: 'Index',
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
