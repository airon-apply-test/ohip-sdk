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
exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerToJSON = exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSONTyped = exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSON = exports.instanceOfChannelAccountContractCopyResponseTypeContractResponseInformationInner = void 0;
const runtime_1 = require("../runtime");
const ErrorType_1 = require("./ErrorType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ChannelAccountContractCopyResponseTypeContractResponseInformationInner interface.
 */
function instanceOfChannelAccountContractCopyResponseTypeContractResponseInformationInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelAccountContractCopyResponseTypeContractResponseInformationInner = instanceOfChannelAccountContractCopyResponseTypeContractResponseInformationInner;
function ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSON(json) {
    return ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSONTyped(json, false);
}
exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSON = ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSON;
function ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'contractId': !(0, runtime_1.exists)(json, 'contractId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['contractId']),
        'sourceContractNo': !(0, runtime_1.exists)(json, 'sourceContractNo') ? undefined : json['sourceContractNo'],
        'targetContractNo': !(0, runtime_1.exists)(json, 'targetContractNo') ? undefined : json['targetContractNo'],
        'errorMessage': !(0, runtime_1.exists)(json, 'errorMessage') ? undefined : (0, ErrorType_1.ErrorTypeFromJSON)(json['errorMessage']),
        'success': !(0, runtime_1.exists)(json, 'success') ? undefined : json['success'],
    };
}
exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSONTyped = ChannelAccountContractCopyResponseTypeContractResponseInformationInnerFromJSONTyped;
function ChannelAccountContractCopyResponseTypeContractResponseInformationInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'contractId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.contractId),
        'sourceContractNo': value.sourceContractNo,
        'targetContractNo': value.targetContractNo,
        'errorMessage': (0, ErrorType_1.ErrorTypeToJSON)(value.errorMessage),
        'success': value.success,
    };
}
exports.ChannelAccountContractCopyResponseTypeContractResponseInformationInnerToJSON = ChannelAccountContractCopyResponseTypeContractResponseInformationInnerToJSON;
