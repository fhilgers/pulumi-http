import * as http from "@pulumi/http";
import * as pulumi from "@pulumi/pulumi";

const tlsConfig = new pulumi.Config('tls');
const vaultConfig = new pulumi.Config('vault');
const tailscaleConfig = new pulumi.Config('tailscale');

const tailnet = tailscaleConfig.requireSecret('tailnet')
const apiKey = tailscaleConfig.requireSecret('apikey')
const token = vaultConfig.getSecret('token')

const vaultInit = new http.Request('vault-init', {
    create: {
        method: 'PUT',
        url: 'https://127.0.0.1:8200/v1/sys/init',
        body: JSON.stringify({
            recovery_shares: 5,
            recovery_threshold: 3,
        }),
        maxRetries: 100,
        certificates: [ tlsConfig.requireSecret('bundle') ],
        rootCAs: [ tlsConfig.require('ca')],
        retryWaitMin: 1,
        retryWaitMax: 10,
    },
}, { additionalSecretOutputs: [ 'response' ] }).response.body.apply(JSON.parse)

const vaultDelete = new http.Request('vault-delete', {
    delete: {
        method: 'PUT',
        url: 'https://127.0.0.1:8200/v1/sys/seal',
        maxRetries: 100,
        header: {
            'X-Vault-Token': [ vaultInit.apply(b => b as { root_token: string }).root_token ]
        },
        certificates: [ tlsConfig.requireSecret('bundle') ],
        rootCAs: [ tlsConfig.require('ca')],
        expectedStatusCode: 204,
        retryWaitMin: 1,
        retryWaitMax: 10,
    }
})


 //const vaultInit = new http.Request('vault-init', {
 //    create: request,
 //    delete: deleteRequest
 //    triggers: [ x ],
 //}, { additionalSecretOutputs: [ 'responsebody' ] } ).responsebody.apply(JSON.parse);


//const devices = new http.Request('tailscale', {
//    method: 'GET',
//    url: pulumi.interpolate`https://api.tailscale.com/api/v2/tailnet/${tailnet}/devices?all`,
//    expectedstatus: 200,
//    headers: {
//        'Authorization': pulumi.interpolate`Basic ${auth}`
//    },
//}).responsebody.apply(JSON.parse)

//const consulInit = new http.Request('consul-init', {
//    method: 'PUT',
//    url: 'http://127.0.0.1:8500/v1/acl/bootstrap',
//}).responsebody.apply(JSON.parse)
//
//const nomadInit = new http.Request('nomad-init', {
//    method: 'POST',
//    url: 'http://127.0.0.1:4646/v1/acl/bootstrap',
//}).responsebody.apply(JSON.parse)

export = async () => {
    return {
        vault: vaultInit,
 //      consul: consulInit,
 //      nomad: nomadInit,
    }
}
