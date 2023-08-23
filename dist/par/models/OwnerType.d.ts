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
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { EmailInfoType } from './EmailInfoType';
import type { PersonNameType } from './PersonNameType';
import type { ProfileId } from './ProfileId';
import type { TelephoneInfoType } from './TelephoneInfoType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Generic type for information about an owner.
 * @export
 * @interface OwnerType
 */
export interface OwnerType {
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof OwnerType
     */
    hotel?: CodeDescriptionType;
    /**
     *
     * @type {UniqueIDType}
     * @memberof OwnerType
     */
    userId?: UniqueIDType;
    /**
     * Unique application user name of the owner.
     * @type {string}
     * @memberof OwnerType
     */
    userName?: string;
    /**
     * Unique Code to identify the owner.
     * @type {string}
     * @memberof OwnerType
     */
    ownerCode?: string;
    /**
     *
     * @type {ProfileId}
     * @memberof OwnerType
     */
    profileId?: ProfileId;
    /**
     *
     * @type {PersonNameType}
     * @memberof OwnerType
     */
    name?: PersonNameType;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof OwnerType
     */
    department?: CodeDescriptionType;
    /**
     *
     * @type {EmailInfoType}
     * @memberof OwnerType
     */
    email?: EmailInfoType;
    /**
     *
     * @type {TelephoneInfoType}
     * @memberof OwnerType
     */
    phone?: TelephoneInfoType;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof OwnerType
     */
    relationship?: CodeDescriptionType;
    /**
     * When true, this is a primary owner.
     * @type {boolean}
     * @memberof OwnerType
     */
    primary?: boolean;
}
/**
 * Check if a given object implements the OwnerType interface.
 */
export declare function instanceOfOwnerType(value: object): boolean;
export declare function OwnerTypeFromJSON(json: any): OwnerType;
export declare function OwnerTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): OwnerType;
export declare function OwnerTypeToJSON(value?: OwnerType | null): any;
