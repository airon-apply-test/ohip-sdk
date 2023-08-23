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
import type { AddressInfoType } from './AddressInfoType';
import type { CompanyType } from './CompanyType';
import type { CustomerType } from './CustomerType';
import type { EmailInfoType } from './EmailInfoType';
import type { OwnerType } from './OwnerType';
import type { ProfileStatusType } from './ProfileStatusType';
import type { ProfileTypeType } from './ProfileTypeType';
import type { TelephoneInfoType } from './TelephoneInfoType';
import type { URLInfoType } from './URLInfoType';
import type { UniqueIDType } from './UniqueIDType';
/**
 *
 * @export
 * @interface RelationshipProfileType
 */
export interface RelationshipProfileType {
    /**
     *
     * @type {CustomerType}
     * @memberof RelationshipProfileType
     */
    customer?: CustomerType;
    /**
     *
     * @type {CompanyType}
     * @memberof RelationshipProfileType
     */
    company?: CompanyType;
    /**
     *
     * @type {TelephoneInfoType}
     * @memberof RelationshipProfileType
     */
    telephone?: TelephoneInfoType;
    /**
     *
     * @type {AddressInfoType}
     * @memberof RelationshipProfileType
     */
    address?: AddressInfoType;
    /**
     *
     * @type {EmailInfoType}
     * @memberof RelationshipProfileType
     */
    email?: EmailInfoType;
    /**
     *
     * @type {URLInfoType}
     * @memberof RelationshipProfileType
     */
    uRLs?: URLInfoType;
    /**
     *
     * @type {OwnerType}
     * @memberof RelationshipProfileType
     */
    primaryOwner?: OwnerType;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RelationshipProfileType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * Indicates if this relationship is the primary relationship.
     * @type {string}
     * @memberof RelationshipProfileType
     */
    primary?: string;
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof RelationshipProfileType
     */
    id?: string;
    /**
     *
     * @type {ProfileStatusType}
     * @memberof RelationshipProfileType
     */
    statusCode?: ProfileStatusType;
    /**
     *
     * @type {ProfileTypeType}
     * @memberof RelationshipProfileType
     */
    profileType?: ProfileTypeType;
}
/**
 * Check if a given object implements the RelationshipProfileType interface.
 */
export declare function instanceOfRelationshipProfileType(value: object): boolean;
export declare function RelationshipProfileTypeFromJSON(json: any): RelationshipProfileType;
export declare function RelationshipProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipProfileType;
export declare function RelationshipProfileTypeToJSON(value?: RelationshipProfileType | null): any;
