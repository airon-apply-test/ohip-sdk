"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PutBlockRefusedReasonsRequestToJSON = exports.PutBlockRefusedReasonsRequestFromJSONTyped = exports.PutBlockRefusedReasonsRequestFromJSON = exports.instanceOfPutBlockRefusedReasonsRequest = void 0;
const runtime_1 = require("../runtime");
const BlockRefusedReasonType_1 = require("./BlockRefusedReasonType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PutBlockRefusedReasonsRequest interface.
 */
function instanceOfPutBlockRefusedReasonsRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPutBlockRefusedReasonsRequest = instanceOfPutBlockRefusedReasonsRequest;
function PutBlockRefusedReasonsRequestFromJSON(json) {
    return PutBlockRefusedReasonsRequestFromJSONTyped(json, false);
}
exports.PutBlockRefusedReasonsRequestFromJSON = PutBlockRefusedReasonsRequestFromJSON;
function PutBlockRefusedReasonsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'blockRefusedReasons': !(0, runtime_1.exists)(json, 'blockRefusedReasons') ? undefined : (json['blockRefusedReasons'].map(BlockRefusedReasonType_1.BlockRefusedReasonTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PutBlockRefusedReasonsRequestFromJSONTyped = PutBlockRefusedReasonsRequestFromJSONTyped;
function PutBlockRefusedReasonsRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'blockRefusedReasons': value.blockRefusedReasons === undefined ? undefined : (value.blockRefusedReasons.map(BlockRefusedReasonType_1.BlockRefusedReasonTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PutBlockRefusedReasonsRequestToJSON = PutBlockRefusedReasonsRequestToJSON;
