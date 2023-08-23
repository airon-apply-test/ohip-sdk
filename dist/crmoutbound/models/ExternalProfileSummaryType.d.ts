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
import type { AddressInfoType } from './AddressInfoType';
import type { EmailInfoType } from './EmailInfoType';
import type { ExternalProfileSummaryTypeFormerName } from './ExternalProfileSummaryTypeFormerName';
import type { OwnerType } from './OwnerType';
import type { ProfileMembershipType } from './ProfileMembershipType';
import type { ProfileStatusType } from './ProfileStatusType';
import type { ProfileTypeType } from './ProfileTypeType';
import type { TelephoneInfoType } from './TelephoneInfoType';
import type { URLInfoType } from './URLInfoType';
/**
 * Type provides the basic information about the external profile.
 * @export
 * @interface ExternalProfileSummaryType
 */
export interface ExternalProfileSummaryType {
    /**
     *
     * @type {ExternalProfileSummaryTypeFormerName}
     * @memberof ExternalProfileSummaryType
     */
    formerName?: ExternalProfileSummaryTypeFormerName;
    /**
     *
     * @type {AddressInfoType}
     * @memberof ExternalProfileSummaryType
     */
    addressInfo?: AddressInfoType;
    /**
     *
     * @type {TelephoneInfoType}
     * @memberof ExternalProfileSummaryType
     */
    telephoneInfo?: TelephoneInfoType;
    /**
     *
     * @type {EmailInfoType}
     * @memberof ExternalProfileSummaryType
     */
    emailInfo?: EmailInfoType;
    /**
     *
     * @type {ProfileMembershipType}
     * @memberof ExternalProfileSummaryType
     */
    profileMembership?: ProfileMembershipType;
    /**
     *
     * @type {URLInfoType}
     * @memberof ExternalProfileSummaryType
     */
    urlInfo?: URLInfoType;
    /**
     * Generic type for a list of owners.
     * @type {Array<OwnerType>}
     * @memberof ExternalProfileSummaryType
     */
    owners?: Array<OwnerType>;
    /**
     *
     * @type {ProfileTypeType}
     * @memberof ExternalProfileSummaryType
     */
    profileType?: ProfileTypeType;
    /**
     *
     * @type {ProfileStatusType}
     * @memberof ExternalProfileSummaryType
     */
    statusCode?: ProfileStatusType;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof ExternalProfileSummaryType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof ExternalProfileSummaryType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof ExternalProfileSummaryType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof ExternalProfileSummaryType
     */
    lastModifierId?: string;
    /**
     * Property this profile is registered with.
     * @type {string}
     * @memberof ExternalProfileSummaryType
     */
    registeredProperty?: string;
}
/**
 * Check if a given object implements the ExternalProfileSummaryType interface.
 */
export declare function instanceOfExternalProfileSummaryType(value: object): boolean;
export declare function ExternalProfileSummaryTypeFromJSON(json: any): ExternalProfileSummaryType;
export declare function ExternalProfileSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExternalProfileSummaryType;
export declare function ExternalProfileSummaryTypeToJSON(value?: ExternalProfileSummaryType | null): any;
