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
exports.NextBlockStatusCodeDetailsToJSON = exports.NextBlockStatusCodeDetailsFromJSONTyped = exports.NextBlockStatusCodeDetailsFromJSON = exports.instanceOfNextBlockStatusCodeDetails = void 0;
const runtime_1 = require("../runtime");
const BlockStatusCodeType_1 = require("./BlockStatusCodeType");
const InstanceLink_1 = require("./InstanceLink");
const NextBlockStatusCodeType_1 = require("./NextBlockStatusCodeType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the NextBlockStatusCodeDetails interface.
 */
function instanceOfNextBlockStatusCodeDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfNextBlockStatusCodeDetails = instanceOfNextBlockStatusCodeDetails;
function NextBlockStatusCodeDetailsFromJSON(json) {
    return NextBlockStatusCodeDetailsFromJSONTyped(json, false);
}
exports.NextBlockStatusCodeDetailsFromJSON = NextBlockStatusCodeDetailsFromJSON;
function NextBlockStatusCodeDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'nextBlockStatusCodeList': !(0, runtime_1.exists)(json, 'nextBlockStatusCodeList') ? undefined : (json['nextBlockStatusCodeList'].map(NextBlockStatusCodeType_1.NextBlockStatusCodeTypeFromJSON)),
        'blockStatusCodeMasterList': !(0, runtime_1.exists)(json, 'blockStatusCodeMasterList') ? undefined : (json['blockStatusCodeMasterList'].map(BlockStatusCodeType_1.BlockStatusCodeTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.NextBlockStatusCodeDetailsFromJSONTyped = NextBlockStatusCodeDetailsFromJSONTyped;
function NextBlockStatusCodeDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'nextBlockStatusCodeList': value.nextBlockStatusCodeList === undefined ? undefined : (value.nextBlockStatusCodeList.map(NextBlockStatusCodeType_1.NextBlockStatusCodeTypeToJSON)),
        'blockStatusCodeMasterList': value.blockStatusCodeMasterList === undefined ? undefined : (value.blockStatusCodeMasterList.map(BlockStatusCodeType_1.BlockStatusCodeTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.NextBlockStatusCodeDetailsToJSON = NextBlockStatusCodeDetailsToJSON;
