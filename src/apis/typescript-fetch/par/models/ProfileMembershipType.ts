/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BenefitType } from './BenefitType';
import {
    BenefitTypeFromJSON,
    BenefitTypeFromJSONTyped,
    BenefitTypeToJSON,
} from './BenefitType';
import type { CardReIssueType } from './CardReIssueType';
import {
    CardReIssueTypeFromJSON,
    CardReIssueTypeFromJSONTyped,
    CardReIssueTypeToJSON,
} from './CardReIssueType';
import type { DowngradeType } from './DowngradeType';
import {
    DowngradeTypeFromJSON,
    DowngradeTypeFromJSONTyped,
    DowngradeTypeToJSON,
} from './DowngradeType';
import type { MembershipEarningPreferenceType } from './MembershipEarningPreferenceType';
import {
    MembershipEarningPreferenceTypeFromJSON,
    MembershipEarningPreferenceTypeFromJSONTyped,
    MembershipEarningPreferenceTypeToJSON,
} from './MembershipEarningPreferenceType';
import type { ParagraphType } from './ParagraphType';
import {
    ParagraphTypeFromJSON,
    ParagraphTypeFromJSONTyped,
    ParagraphTypeToJSON,
} from './ParagraphType';
import type { TierAdministrationType } from './TierAdministrationType';
import {
    TierAdministrationTypeFromJSON,
    TierAdministrationTypeFromJSONTyped,
    TierAdministrationTypeToJSON,
} from './TierAdministrationType';

/**
 * Detailed information of the memberships.
 * @export
 * @interface ProfileMembershipType
 */
export interface ProfileMembershipType {
    /**
     * 
     * @type {ParagraphType}
     * @memberof ProfileMembershipType
     */
    comment?: ParagraphType;
    /**
     * Card Number of the membership.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    newMembershipNumber?: string;
    /**
     * Name to be displayed on the membership card.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    nameOnCard?: string;
    /**
     * Description of the membership program.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    programDescription?: string;
    /**
     * Indicates the membership level.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipLevel?: string;
    /**
     * Indicates the membership level description.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipLevelDescription?: string;
    /**
     * Indicates the membership class.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipClass?: string;
    /**
     * 
     * @type {MembershipEarningPreferenceType}
     * @memberof ProfileMembershipType
     */
    earningPreference?: MembershipEarningPreferenceType;
    /**
     * Indicates whether membership is active or inactive.
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    inactive?: boolean;
    /**
     * Basic information about membership benefit.
     * @type {Array<BenefitType>}
     * @memberof ProfileMembershipType
     */
    benefits?: Array<BenefitType>;
    /**
     * 
     * @type {TierAdministrationType}
     * @memberof ProfileMembershipType
     */
    tierAdministration?: TierAdministrationType;
    /**
     * 
     * @type {DowngradeType}
     * @memberof ProfileMembershipType
     */
    downgrade?: DowngradeType;
    /**
     * 
     * @type {CardReIssueType}
     * @memberof ProfileMembershipType
     */
    reIssueNewCard?: CardReIssueType;
    /**
     * True if you want to exclude the member from the Membership Fulfillment extract,the member's actions will not be included in the fulfillment extract until this value set to false.
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    excludeFromBatch?: boolean;
    /**
     * Indicates Upgrade information which includes member's next tier level, policyRequirements for the next upgrade.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    upgradeDescription?: string;
    /**
     * Indicates information regarding the member's possible downgrades.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    downgradeDescription?: string;
    /**
     * Value Rating Type Description for this membership.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    rating?: string;
    /**
     * Indicates how the guest enrolled in the program.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipEnrollmentCode?: string;
    /**
     * Indicates where the guest is in the membership enrollment process.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    memberStatus?: string;
    /**
     * Profile MemberShip Points.
     * @type {number}
     * @memberof ProfileMembershipType
     */
    currentPoints?: number;
    /**
     * Label used to refer to points for this membership type
     * @type {string}
     * @memberof ProfileMembershipType
     */
    pointsLabel?: string;
    /**
     * Source from where the enrollment is done.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    enrollmentSource?: string;
    /**
     * Resort/CRO where enrollment is done.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    enrollmentResort?: string;
    /**
     * Preferred Card.
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    preferredCard?: boolean;
    /**
     * Card Number of the membership.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipId?: string;
    /**
     * Type of membership.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    membershipType?: string;
    /**
     * Indicator if Membership is a Primary Membership.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    primaryMembershipYn?: string;
    /**
     * Boolean indicator set to True implies membership is a Primary Membership.
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    primaryMembership?: boolean;
    /**
     * Membership ID Number.
     * @type {number}
     * @memberof ProfileMembershipType
     */
    membershipIdNo?: number;
    /**
     * Ranking assigned to the Player Profile by the Gaming system.
     * @type {number}
     * @memberof ProfileMembershipType
     */
    playerRanking?: number;
    /**
     * Indicates how the award points for this membership type will be managed.
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    centralSetup?: boolean;
    /**
     * Indicates when the member signed up for the loyalty program.
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    signupDate?: Date;
    /**
     * Indicates the starting date.
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    effectiveDate?: Date;
    /**
     * Indicates the ending date.
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    expireDate?: Date;
    /**
     * When true, indicates that the ExpireDate is the first day after the applicable period (e.g. when expire date is Oct 15 the last date of the period is Oct 14).
     * @type {boolean}
     * @memberof ProfileMembershipType
     */
    expireDateExclusiveIndicator?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ProfileMembershipType
     */
    orderSequence?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof ProfileMembershipType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof ProfileMembershipType
     */
    purgeDate?: Date;
}

/**
 * Check if a given object implements the ProfileMembershipType interface.
 */
export function instanceOfProfileMembershipType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileMembershipTypeFromJSON(json: any): ProfileMembershipType {
    return ProfileMembershipTypeFromJSONTyped(json, false);
}

export function ProfileMembershipTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileMembershipType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'comment': !exists(json, 'comment') ? undefined : ParagraphTypeFromJSON(json['comment']),
        'newMembershipNumber': !exists(json, 'newMembershipNumber') ? undefined : json['newMembershipNumber'],
        'nameOnCard': !exists(json, 'nameOnCard') ? undefined : json['nameOnCard'],
        'programDescription': !exists(json, 'programDescription') ? undefined : json['programDescription'],
        'membershipLevel': !exists(json, 'membershipLevel') ? undefined : json['membershipLevel'],
        'membershipLevelDescription': !exists(json, 'membershipLevelDescription') ? undefined : json['membershipLevelDescription'],
        'membershipClass': !exists(json, 'membershipClass') ? undefined : json['membershipClass'],
        'earningPreference': !exists(json, 'earningPreference') ? undefined : MembershipEarningPreferenceTypeFromJSON(json['earningPreference']),
        'inactive': !exists(json, 'inactive') ? undefined : json['inactive'],
        'benefits': !exists(json, 'benefits') ? undefined : ((json['benefits'] as Array<any>).map(BenefitTypeFromJSON)),
        'tierAdministration': !exists(json, 'tierAdministration') ? undefined : TierAdministrationTypeFromJSON(json['tierAdministration']),
        'downgrade': !exists(json, 'downgrade') ? undefined : DowngradeTypeFromJSON(json['downgrade']),
        'reIssueNewCard': !exists(json, 'reIssueNewCard') ? undefined : CardReIssueTypeFromJSON(json['reIssueNewCard']),
        'excludeFromBatch': !exists(json, 'excludeFromBatch') ? undefined : json['excludeFromBatch'],
        'upgradeDescription': !exists(json, 'upgradeDescription') ? undefined : json['upgradeDescription'],
        'downgradeDescription': !exists(json, 'downgradeDescription') ? undefined : json['downgradeDescription'],
        'rating': !exists(json, 'rating') ? undefined : json['rating'],
        'membershipEnrollmentCode': !exists(json, 'membershipEnrollmentCode') ? undefined : json['membershipEnrollmentCode'],
        'memberStatus': !exists(json, 'memberStatus') ? undefined : json['memberStatus'],
        'currentPoints': !exists(json, 'currentPoints') ? undefined : json['currentPoints'],
        'pointsLabel': !exists(json, 'pointsLabel') ? undefined : json['pointsLabel'],
        'enrollmentSource': !exists(json, 'enrollmentSource') ? undefined : json['enrollmentSource'],
        'enrollmentResort': !exists(json, 'enrollmentResort') ? undefined : json['enrollmentResort'],
        'preferredCard': !exists(json, 'preferredCard') ? undefined : json['preferredCard'],
        'membershipId': !exists(json, 'membershipId') ? undefined : json['membershipId'],
        'membershipType': !exists(json, 'membershipType') ? undefined : json['membershipType'],
        'primaryMembershipYn': !exists(json, 'primaryMembershipYn') ? undefined : json['primaryMembershipYn'],
        'primaryMembership': !exists(json, 'primaryMembership') ? undefined : json['primaryMembership'],
        'membershipIdNo': !exists(json, 'membershipIdNo') ? undefined : json['membershipIdNo'],
        'playerRanking': !exists(json, 'playerRanking') ? undefined : json['playerRanking'],
        'centralSetup': !exists(json, 'centralSetup') ? undefined : json['centralSetup'],
        'signupDate': !exists(json, 'signupDate') ? undefined : (new Date(json['signupDate'])),
        'effectiveDate': !exists(json, 'effectiveDate') ? undefined : (new Date(json['effectiveDate'])),
        'expireDate': !exists(json, 'expireDate') ? undefined : (new Date(json['expireDate'])),
        'expireDateExclusiveIndicator': !exists(json, 'expireDateExclusiveIndicator') ? undefined : json['expireDateExclusiveIndicator'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}

export function ProfileMembershipTypeToJSON(value?: ProfileMembershipType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'comment': ParagraphTypeToJSON(value.comment),
        'newMembershipNumber': value.newMembershipNumber,
        'nameOnCard': value.nameOnCard,
        'programDescription': value.programDescription,
        'membershipLevel': value.membershipLevel,
        'membershipLevelDescription': value.membershipLevelDescription,
        'membershipClass': value.membershipClass,
        'earningPreference': MembershipEarningPreferenceTypeToJSON(value.earningPreference),
        'inactive': value.inactive,
        'benefits': value.benefits === undefined ? undefined : ((value.benefits as Array<any>).map(BenefitTypeToJSON)),
        'tierAdministration': TierAdministrationTypeToJSON(value.tierAdministration),
        'downgrade': DowngradeTypeToJSON(value.downgrade),
        'reIssueNewCard': CardReIssueTypeToJSON(value.reIssueNewCard),
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
        'signupDate': value.signupDate === undefined ? undefined : (value.signupDate.toISOString().substr(0,10)),
        'effectiveDate': value.effectiveDate === undefined ? undefined : (value.effectiveDate.toISOString().substr(0,10)),
        'expireDate': value.expireDate === undefined ? undefined : (value.expireDate.toISOString().substr(0,10)),
        'expireDateExclusiveIndicator': value.expireDateExclusiveIndicator,
        'orderSequence': value.orderSequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
    };
}

