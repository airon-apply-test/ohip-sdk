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
import type { BenefitType } from './BenefitType';
import type { CardReIssueType } from './CardReIssueType';
import type { DowngradeType } from './DowngradeType';
import type { MembershipEarningPreferenceType } from './MembershipEarningPreferenceType';
import type { ParagraphType } from './ParagraphType';
import type { TierAdministrationType } from './TierAdministrationType';
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
     * Indicates Upgrade information which includes member's next tier level, requirements for the next upgrade.
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
export declare function instanceOfProfileMembershipType(value: object): boolean;
export declare function ProfileMembershipTypeFromJSON(json: any): ProfileMembershipType;
export declare function ProfileMembershipTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileMembershipType;
export declare function ProfileMembershipTypeToJSON(value?: ProfileMembershipType | null): any;
