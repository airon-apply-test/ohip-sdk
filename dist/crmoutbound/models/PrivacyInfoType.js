"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PrivacyInfoTypeToJSON = exports.PrivacyInfoTypeFromJSONTyped = exports.PrivacyInfoTypeFromJSON = exports.instanceOfPrivacyInfoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the PrivacyInfoType interface.
 */
function instanceOfPrivacyInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPrivacyInfoType = instanceOfPrivacyInfoType;
function PrivacyInfoTypeFromJSON(json) {
    return PrivacyInfoTypeFromJSONTyped(json, false);
}
exports.PrivacyInfoTypeFromJSON = PrivacyInfoTypeFromJSON;
function PrivacyInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'marketResearchParticipation': !(0, runtime_1.exists)(json, 'marketResearchParticipation') ? undefined : json['marketResearchParticipation'],
        'lastPrivacyPromptDate': !(0, runtime_1.exists)(json, 'lastPrivacyPromptDate') ? undefined : (new Date(json['lastPrivacyPromptDate'])),
        'infoFromThirdParty': !(0, runtime_1.exists)(json, 'infoFromThirdParty') ? undefined : json['infoFromThirdParty'],
        'autoEnrollLoyaltyProgram': !(0, runtime_1.exists)(json, 'autoEnrollLoyaltyProgram') ? undefined : json['autoEnrollLoyaltyProgram'],
        'allowPhone': !(0, runtime_1.exists)(json, 'allowPhone') ? undefined : json['allowPhone'],
        'allowSMS': !(0, runtime_1.exists)(json, 'allowSMS') ? undefined : json['allowSMS'],
        'allowEmail': !(0, runtime_1.exists)(json, 'allowEmail') ? undefined : json['allowEmail'],
        'optInMailingList': !(0, runtime_1.exists)(json, 'optInMailingList') ? undefined : json['optInMailingList'],
        'optInMarketResearch': !(0, runtime_1.exists)(json, 'optInMarketResearch') ? undefined : json['optInMarketResearch'],
        'optInThirdParty': !(0, runtime_1.exists)(json, 'optInThirdParty') ? undefined : json['optInThirdParty'],
        'optInAutoEnrollmentMember': !(0, runtime_1.exists)(json, 'optInAutoEnrollmentMember') ? undefined : json['optInAutoEnrollmentMember'],
        'optInPhone': !(0, runtime_1.exists)(json, 'optInPhone') ? undefined : json['optInPhone'],
        'optInSms': !(0, runtime_1.exists)(json, 'optInSms') ? undefined : json['optInSms'],
        'optInEmail': !(0, runtime_1.exists)(json, 'optInEmail') ? undefined : json['optInEmail'],
    };
}
exports.PrivacyInfoTypeFromJSONTyped = PrivacyInfoTypeFromJSONTyped;
function PrivacyInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'marketResearchParticipation': value.marketResearchParticipation,
        'lastPrivacyPromptDate': value.lastPrivacyPromptDate === undefined ? undefined : (value.lastPrivacyPromptDate.toISOString().substr(0, 10)),
        'infoFromThirdParty': value.infoFromThirdParty,
        'autoEnrollLoyaltyProgram': value.autoEnrollLoyaltyProgram,
        'allowPhone': value.allowPhone,
        'allowSMS': value.allowSMS,
        'allowEmail': value.allowEmail,
        'optInMailingList': value.optInMailingList,
        'optInMarketResearch': value.optInMarketResearch,
        'optInThirdParty': value.optInThirdParty,
        'optInAutoEnrollmentMember': value.optInAutoEnrollmentMember,
        'optInPhone': value.optInPhone,
        'optInSms': value.optInSms,
        'optInEmail': value.optInEmail,
    };
}
exports.PrivacyInfoTypeToJSON = PrivacyInfoTypeToJSON;
