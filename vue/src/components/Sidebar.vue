<template>
	<SpSidebar
		v-on:sidebar-open="sidebarOpen = true"
		v-on:sidebar-close="sidebarOpen = false"
	>
		<template v-slot:default>
			<SpLinkIcon link="/" text="Dashboard" icon="Dashboard" />
			<SpLinkIcon link="/types" text="NameService" icon="Form" />
			<SpLinkIcon link="/relayers" text="Gravity DEX" icon="Transactions" />
			<div class="sp-dash"></div>
			<SpLinkIcon
				href="https://github.com/Orionis-Labs/orionis"
				target="_blank"
				text="Source Code"
				icon="GitHub"
			/>
		</template>
		<template v-slot:footer>
			<div class="sp-text">Build: v0.3.8</div>
		</template>
	</SpSidebar>
</template>

<script>
export default {
	name: 'Sidebar',
	data() {
		return {
			sidebarOpen: true
		}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	}
}
</script>
