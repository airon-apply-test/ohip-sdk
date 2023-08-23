"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationMembershipTypeToJSON = exports.ReservationMembershipTypeFromJSONTyped = exports.ReservationMembershipTypeFromJSON = exports.instanceOfReservationMembershipType = void 0;
const runtime_1 = require("../runtime");
const BenefitType_1 = require("./BenefitType");
const CardReIssueType_1 = require("./CardReIssueType");
const DowngradeType_1 = require("./DowngradeType");
const MembershipEarningPreferenceType_1 = require("./MembershipEarningPreferenceType");
const ParagraphType_1 = require("./ParagraphType");
const TierAdministrationType_1 = require("./TierAdministrationType");
/**
 * Check if a given object implements the ReservationMembershipType interface.
 */
function instanceOfReservationMembershipType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationMembershipType = instanceOfReservationMembershipType;
function ReservationMembershipTypeFromJSON(json) {
    return ReservationMembershipTypeFromJSONTyped(json, false);
}
exports.ReservationMembershipTypeFromJSON = ReservationMembershipTypeFromJSON;
function ReservationMembershipTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'comment': !(0, runtime_1.exists)(json, 'comment') ? undefined : (0, ParagraphType_1.ParagraphTypeFromJSON)(json['comment']),
        'newMembershipNumber': !(0, runtime_1.exists)(json, 'newMembershipNumber') ? undefined : json['newMembershipNumber'],
        'nameOnCard': !(0, runtime_1.exists)(json, 'nameOnCard') ? undefined : json['nameOnCard'],
        'programDescription': !(0, runtime_1.exists)(json, 'programDescription') ? undefined : json['programDescription'],
        'membershipLevel': !(0, runtime_1.exists)(json, 'membershipLevel') ? undefined : json['membershipLevel'],
        'membershipLevelDescription': !(0, runtime_1.exists)(json, 'membershipLevelDescription') ? undefined : json['membershipLevelDescription'],
        'membershipClass': !(0, runtime_1.exists)(json, 'membershipClass') ? undefined : json['membershipClass'],
        'earningPreference': !(0, runtime_1.exists)(json, 'earningPreference') ? undefined : (0, MembershipEarningPreferenceType_1.MembershipEarningPreferenceTypeFromJSON)(json['earningPreference']),
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
        'benefits': !(0, runtime_1.exists)(json, 'benefits') ? undefined : (json['benefits'].map(BenefitType_1.BenefitTypeFromJSON)),
        'tierAdministration': !(0, runtime_1.exists)(json, 'tierAdministration') ? undefined : (0, TierAdministrationType_1.TierAdministrationTypeFromJSON)(json['tierAdministration']),
        'downgrade': !(0, runtime_1.exists)(json, 'downgrade') ? undefined : (0, DowngradeType_1.DowngradeTypeFromJSON)(json['downgrade']),
        'reIssueNewCard': !(0, runtime_1.exists)(json, 'reIssueNewCard') ? undefined : (0, CardReIssueType_1.CardReIssueTypeFromJSON)(json['reIssueNewCard']),
        'excludeFromBatch': !(0, runtime_1.exists)(json, 'excludeFromBatch') ? undefined : json['excludeFromBatch'],
        'upgradeDescription': !(0, runtime_1.exists)(json, 'upgradeDescription') ? undefined : json['upgradeDescription'],
        'downgradeDescription': !(0, runtime_1.exists)(json, 'downgradeDescription') ? undefined : json['downgradeDescription'],
        'rating': !(0, runtime_1.exists)(json, 'rating') ? undefined : json['rating'],
        'membershipEnrollmentCode': !(0, runtime_1.exists)(json, 'membershipEnrollmentCode') ? undefined : json['membershipEnrollmentCode'],
        'memberStatus': !(0, runtime_1.exists)(json, 'memberStatus') ? undefined : json['memberStatus'],
        'currentPoints': !(0, runtime_1.exists)(json, 'currentPoints') ? undefined : json['currentPoints'],
        'pointsLabel': !(0, runtime_1.exists)(json, 'pointsLabel') ? undefined : json['pointsLabel'],
        'enrollmentSource': !(0, runtime_1.exists)(json, 'enrollmentSource') ? undefined : json['enrollmentSource'],
        'enrollmentResort': !(0, runtime_1.exists)(json, 'enrollmentResort') ? undefined : json['enrollmentResort'],
        'preferredCard': !(0, runtime_1.exists)(json, 'preferredCard') ? undefined : json['preferredCard'],
        'membershipId': !(0, runtime_1.exists)(json, 'membershipId') ? undefined : json['membershipId'],
        'membershipType': !(0, runtime_1.exists)(json, 'membershipType') ? undefined : json['membershipType'],
        'primaryMembershipYn': !(0, runtime_1.exists)(json, 'primaryMembershipYn') ? undefined : json['primaryMembershipYn'],
        'primaryMembership': !(0, runtime_1.exists)(json, 'primaryMembership') ? undefined : json['primaryMembership'],
        'membershipIdNo': !(0, runtime_1.exists)(json, 'membershipIdNo') ? undefined : json['membershipIdNo'],
        'playerRanking': !(0, runtime_1.exists)(json, 'playerRanking') ? undefined : json['playerRanking'],
        'centralSetup': !(0, runtime_1.exists)(json, 'centralSetup') ? undefined : json['centralSetup'],
        'signupDate': !(0, runtime_1.exists)(json, 'signupDate') ? undefined : (new Date(json['signupDate'])),
        'effectiveDate': !(0, runtime_1.exists)(json, 'effectiveDate') ? undefined : (new Date(json['effectiveDate'])),
        'expireDate': !(0, runtime_1.exists)(json, 'expireDate') ? undefined : (new Date(json['expireDate'])),
        'expireDateExclusiveIndicator': !(0, runtime_1.exists)(json, 'expireDateExclusiveIndicator') ? undefined : json['expireDateExclusiveIndicator'],
        'orderSequence': !(0, runtime_1.exists)(json, 'orderSequence') ? undefined : json['orderSequence'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'linkMembership': !(0, runtime_1.exists)(json, 'linkMembership') ? undefined : json['linkMembership'],
        'primary': !(0, runtime_1.exists)(json, 'primary') ? undefined : json['primary'],
    };
}
exports.ReservationMembershipTypeFromJSONTyped = ReservationMembershipTypeFromJSONTyped;
function ReservationMembershipTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'comment': (0, ParagraphType_1.ParagraphTypeToJSON)(value.comment),
        'newMembershipNumber': value.newMembershipNumber,
        'nameOnCard': value.nameOnCard,
        'programDescription': value.programDescription,
        'membershipLevel': value.membershipLevel,
        'membershipLevelDescription': value.membershipLevelDescription,
        'membershipClass': value.membershipClass,
        'earningPreference': (0, MembershipEarningPreferenceType_1.MembershipEarningPreferenceTypeToJSON)(value.earningPreference),
        'inactive': value.inactive,
        'benefits': value.benefits === undefined ? undefined : (value.benefits.map(BenefitType_1.BenefitTypeToJSON)),
        'tierAdministration': (0, TierAdministrationType_1.TierAdministrationTypeToJSON)(value.tierAdministration),
        'downgrade': (0, DowngradeType_1.DowngradeTypeToJSON)(value.downgrade),
        'reIssueNewCard': (0, CardReIssueType_1.CardReIssueTypeToJSON)(value.reIssueNewCard),
        'excludeFromBatch': value.excludeFromBatch,
        'upgradeDescription': value.upgradeDescription,
        'downgradeDescription': value.downgradeDescription,
        'rating': value.rating,
        'membershipEnrollmentCode': value.membershipEnrollmentCode,
        'memberStatus': value.memberStatus,
        'currentPoints': value.currentPoints,
        'pointsLabel': value.pointsLabel,
        'enrollmentSource': value.enrollmentSource,
        'enrollmentResort': value.enrollmentResort,
        'preferredCard': value.preferredCard,
        'membershipId': value.membershipId,
        'membershipType': value.membershipType,
        'primaryMembershipYn': value.primaryMembershipYn,
        'primaryMembership': value.primaryMembership,
        'membershipIdNo': value.membershipIdNo,
        'playerRanking': value.playerRanking,
        'centralSetup': value.centralSetup,
        'signupDate': value.signupDate === undefined ? undefined : (value.signupDate.toISOString().substr(0, 10)),
        'effectiveDate': value.effectiveDate === undefined ? undefined : (value.effectiveDate.toISOString().substr(0, 10)),
        'expireDate': value.expireDate === undefined ? undefined : (value.expireDate.toISOString().substr(0, 10)),
        'expireDateExclusiveIndicator': value.expireDateExclusiveIndicator,
        'orderSequence': value.orderSequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'linkMembership': value.linkMembership,
        'primary': value.primary,
    };
}
exports.ReservationMembershipTypeToJSON = ReservationMembershipTypeToJSON;
