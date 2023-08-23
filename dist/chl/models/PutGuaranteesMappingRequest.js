"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PutGuaranteesMappingRequestToJSON = exports.PutGuaranteesMappingRequestFromJSONTyped = exports.PutGuaranteesMappingRequestFromJSON = exports.instanceOfPutGuaranteesMappingRequest = void 0;
const runtime_1 = require("../runtime");
const GuaranteeMappingType_1 = require("./GuaranteeMappingType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PutGuaranteesMappingRequest interface.
 */
function instanceOfPutGuaranteesMappingRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPutGuaranteesMappingRequest = instanceOfPutGuaranteesMappingRequest;
function PutGuaranteesMappingRequestFromJSON(json) {
    return PutGuaranteesMappingRequestFromJSONTyped(json, false);
}
exports.PutGuaranteesMappingRequestFromJSON = PutGuaranteesMappingRequestFromJSON;
function PutGuaranteesMappingRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'guarantees': !(0, runtime_1.exists)(json, 'guarantees') ? undefined : (json['guarantees'].map(GuaranteeMappingType_1.GuaranteeMappingTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PutGuaranteesMappingRequestFromJSONTyped = PutGuaranteesMappingRequestFromJSONTyped;
function PutGuaranteesMappingRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'guarantees': value.guarantees === undefined ? undefined : (value.guarantees.map(GuaranteeMappingType_1.GuaranteeMappingTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PutGuaranteesMappingRequestToJSON = PutGuaranteesMappingRequestToJSON;
