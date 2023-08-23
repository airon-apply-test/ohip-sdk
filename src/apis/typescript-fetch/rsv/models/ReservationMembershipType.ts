/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * @interface ReservationMembershipType
 */
export interface ReservationMembershipType {
    /**
     * 
     * @type {ParagraphType}
     * @memberof ReservationMembershipType
     */
    comment?: ParagraphType;
    /**
     * Card Number of the membership.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    newMembershipNumber?: string;
    /**
     * Name to be displayed on the membership card.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    nameOnCard?: string;
    /**
     * Description of the membership program.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    programDescription?: string;
    /**
     * Indicates the membership level.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipLevel?: string;
    /**
     * Indicates the membership level description.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipLevelDescription?: string;
    /**
     * Indicates the membership class.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipClass?: string;
    /**
     * 
     * @type {MembershipEarningPreferenceType}
     * @memberof ReservationMembershipType
     */
    earningPreference?: MembershipEarningPreferenceType;
    /**
     * Indicates whether membership is active or inactive.
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    inactive?: boolean;
    /**
     * Basic information about membership benefit.
     * @type {Array<BenefitType>}
     * @memberof ReservationMembershipType
     */
    benefits?: Array<BenefitType>;
    /**
     * 
     * @type {TierAdministrationType}
     * @memberof ReservationMembershipType
     */
    tierAdministration?: TierAdministrationType;
    /**
     * 
     * @type {DowngradeType}
     * @memberof ReservationMembershipType
     */
    downgrade?: DowngradeType;
    /**
     * 
     * @type {CardReIssueType}
     * @memberof ReservationMembershipType
     */
    reIssueNewCard?: CardReIssueType;
    /**
     * True if you want to exclude the member from the Membership Fulfillment extract,the member's actions will not be included in the fulfillment extract until this value set to false.
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    excludeFromBatch?: boolean;
    /**
     * Indicates Upgrade information which includes member's next tier level, requirements for the next upgrade.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    upgradeDescription?: string;
    /**
     * Indicates information regarding the member's possible downgrades.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    downgradeDescription?: string;
    /**
     * Value Rating Type Description for this membership.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    rating?: string;
    /**
     * Indicates how the guest enrolled in the program.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipEnrollmentCode?: string;
    /**
     * Indicates where the guest is in the membership enrollment process.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    memberStatus?: string;
    /**
     * Profile MemberShip Points.
     * @type {number}
     * @memberof ReservationMembershipType
     */
    currentPoints?: number;
    /**
     * Label used to refer to points for this membership type
     * @type {string}
     * @memberof ReservationMembershipType
     */
    pointsLabel?: string;
    /**
     * Source from where the enrollment is done.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    enrollmentSource?: string;
    /**
     * Resort/CRO where enrollment is done.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    enrollmentResort?: string;
    /**
     * Preferred Card.
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    preferredCard?: boolean;
    /**
     * Card Number of the membership.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipId?: string;
    /**
     * Type of membership.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    membershipType?: string;
    /**
     * Indicator if Membership is a Primary Membership.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    primaryMembershipYn?: string;
    /**
     * Boolean indicator set to True implies membership is a Primary Membership.
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    primaryMembership?: boolean;
    /**
     * Membership ID Number.
     * @type {number}
     * @memberof ReservationMembershipType
     */
    membershipIdNo?: number;
    /**
     * Ranking assigned to the Player Profile by the Gaming system.
     * @type {number}
     * @memberof ReservationMembershipType
     */
    playerRanking?: number;
    /**
     * Indicates how the award points for this membership type will be managed.
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    centralSetup?: boolean;
    /**
     * Indicates when the member signed up for the loyalty program.
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    signupDate?: Date;
    /**
     * Indicates the starting date.
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    effectiveDate?: Date;
    /**
     * Indicates the ending date.
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    expireDate?: Date;
    /**
     * When true, indicates that the ExpireDate is the first day after the applicable period (e.g. when expire date is Oct 15 the last date of the period is Oct 14).
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    expireDateExclusiveIndicator?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ReservationMembershipType
     */
    orderSequence?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof ReservationMembershipType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof ReservationMembershipType
     */
    purgeDate?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    linkMembership?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof ReservationMembershipType
     */
    primary?: boolean;
}

/**
 * Check if a given object implements the ReservationMembershipType interface.
 */
export function instanceOfReservationMembershipType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationMembershipTypeFromJSON(json: any): ReservationMembershipType {
    return ReservationMembershipTypeFromJSONTyped(json, false);
}

export function ReservationMembershipTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationMembershipType {
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
        'linkMembership': !exists(json, 'linkMembership') ? undefined : json['linkMembership'],
        'primary': !exists(json, 'primary') ? undefined : json['primary'],
    };
}

export function ReservationMembershipTypeToJSON(value?: ReservationMembershipType | null): any {
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
        'linkMembership': value.linkMembership,
        'primary': value.primary,
    };
}

